package auth

import (
	"crypto/md5"
	"errors"
	"math/rand"
	"sync"
	"time"
)

const (
	UNLOCK_TIME time.Duration = 5 * time.Minute
	EXPIRE_TIME time.Duration = 4 * time.Hour
	LOCK_NUM    int32         = 3
	VERIFY_NUM  int32         = 5
	LETTERS     []byte        = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	CODE_LEN    int32         = 4
)

type VerifyInfoHandle interface {
	// 判断是否锁定
	UnAuth() error

	// 判断验证码
	Verify(code string) error

	// 更新验证信息
	Update()

	// 重置验证信息
	Reset()
}

type VerifyInfo struct {
	TokenId    string
	Username   string    // 用户名
	Token      string    // token
	Code       string    // 验证码
	Number     int32     // 登录次数
	CreateTime time.Time // 创建起始时间
	LockTime   time.Time // 锁定起始时间
	vm         *VerifyManager
	// mutex      sync.Mutex
}

type VerifyManager struct {
	VerifyStore sync.Map
}

func NewVerifyManager() *VerifyManager {
	return &VerifyManager{}
}

func (vm *VerifyManager) Run() {
	rand.Seed(time.Now().Unix())
	t := time.NewTicker(10 * time.Second)
	f := func(k, v interface{}) bool {
		interval := time.Now().Sub(v.(*VerifyInfo).LockTime).Seconds()
		if interval > UNLOCK_TIME.Seconds() {
			vm.ResetInfo(v.(*VerifyInfo))
		}
		interval = time.Now().Sub(v.(*VerifyInfo).CreateTime).Seconds()
		if interval > EXPIRE_TIME.Seconds() {
			vm.DelInfo(v.(*VerifyInfo))
		}
		return true
	}

	for {
		select {
		case <-t.C:
			vm.VerifyStore.Range(f)
		}
	}
}

func (vm *VerifyManager) BuildVerifyCode() string {
	code := make([]byte, CODE_LEN)
	lettersLen := len(LETTERS)
	for i := 0; i < int(CODE_LEN); i++ {
		code[i] = LETTERS[rand.Intn(lettersLen)]
	}
	return string(code)
}

func (vm *VerifyManager) GetInfo(name string, token string) *VerifyInfo {
	tokenId := getId([]string{name, token}...)
	vInfo, ok := vm.VerifyStore.Load(tokenId)
	if !ok {
		vInfoP := &VerifyInfo{
			TokenId:    tokenId,
			Username:   name,
			Token:      token,
			Code:       vm.BuildVerifyCode(),
			Number:     0,
			CreateTime: time.Now(),
		}
		vm.VerifyStore.Store(tokenId, vInfoP)
		return vInfoP
	}
	return vInfo.(*VerifyInfo)
}

func (vm *VerifyManager) UpInfo(v *VerifyInfo) {
	v.Number++
	if v.Number >= LOCK_NUM {
		v.LockTime = time.Now()
	}
	vm.VerifyStore.Store(v.TokenId, v)
}

func (vm *VerifyManager) ResetInfo(v *VerifyInfo) {
	v.Number = 0
	vm.VerifyStore.Store(v.TokenId, v)
}

func (vm *VerifyManager) DelInfo(v *VerifyInfo) {
	vm.VerifyStore.Delete(v.TokenId)
}

func (vi *VerifyInfo) UnAuth() error {
	if vi.Number >= LOCK_NUM {
		return errors.New("unauth")
	}
	return nil
}

func (vi *VerifyInfo) Verify(code string) error {
	if vi.Number < VERIFY_NUM {
		return nil
	}
	if code != vi.Code {
		return errors.New("Verify error")
	}

	return nil
}

func (v *VerifyInfo) Update() {
	verifyManagerServer.UpInfo(Update)
}

func (v *VerifyInfo) Reset() {
	vInfo, _ := v.StoreData.Load(name)
	vInfo.Number = 0
	v.StoreData.Store(name)
}

func getId(str ...string) string {
	allstr := ""
	for _, v := range str {
		allstr += v
	}
	tokenId := md5.Sum([]byte(allstr))
	return string(tokenId[:])
}
