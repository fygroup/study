package auth

import (
	"crypto/md5"
	"errors"
	"math/rand"
	"sync"
	"time"
)

var LOCK_USER_TIME time.Duration = 5 * time.Minute
var RESET_TIME time.Duration = 4 * time.Hour
var LOGIN_NUM int32 = 3
var LOCK_NUM int32 = 5
var LETTERS []byte = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var CODELEN int32 = 4

type VSTATE int32

const (
	RESET  VSTATE = 0
	ADD    VSTATE = 1
	DELETE VSTATE = 2
)

type VerifyInf interface {
	// 判断是否锁定
	UnAuth(userName string, token string) bool

	// 判断验证码
	Verify(verifyCode string) bool

	// 更新验证信息
	Update()

	// 重置验证信息
	Reset()
}

type VerifyInfo struct {
	Username   string    // 用户名
	Token      string    // token
	Code       string    // 验证码
	Number     int32     // 登录次数
	CreateTime time.Time // 创建起始时间
	LockTime   time.Time // 锁定起始时间
	// mutex      sync.Mutex
}

type VerifyManager struct {
	VerifyStore sync.Map
	ExpireTime  time.Duration
	UnlockTime  time.Duration
	LockNum     int32
	VerifyNum   int32
	CodeLen     int32
}

func NewVerifyManager() *VerifyManager {
	
} 

func (vm *VerifyManager) Run() {
	rand.Seed(time.Now().Unix())
	t := time.NewTicker(10 * time.Second)
	f := func(k, v interface{}) bool {
		interval := time.Now().Sub(v.(*VerifyInfo).LockTime).Seconds()
		if interval > vm.UnlockTime.Seconds() {
			vm.ResetInfo(v.(*VerifyInfo))
		}
		interval = time.Now().Sub(v.(*VerifyInfo).CreateTime).Seconds()
		if interval > vm.ExpireTime.Seconds() {
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

func (vm *VerifyManager) GetId(str ...string) string {
	allstr := ""
	for _, v := range str {
		allstr += v
	}
	tokenId := md5.Sum([]byte(allstr))
	return string(tokenId[:])
}

func (vm *VerifyManager) GetInfo(name string, token string) *VerifyInfo {
	tokenId := vm.GetId([]string{name, token}...)
	vInfo, ok := vm.VerifyStore.Load(tokenId)
	if !ok {
		vInfoP := &VerifyInfo{
			Username:   name,
			Token:      token,
			Code:       vm.buildVerifyCode(),
			Number:     0,
			CreateTime: time.Now(),
		}
		vm.VerifyStore.Store(tokenId, vInfoP)
		return vInfoP
	}
	return vInfo.(*VerifyInfo)
}

func (vm *VerifyManager) UpInfo(v *VerifyInfo) {
	tokenId := vm.GetId([]string{v.Username, v.Token}...)
	v.Number++
	if v.Number >= vm.LockNum {
		v.LockTime = time.Now()
	}
	vm.VerifyStore.Store(tokenId, v)
}

func (vm *VerifyManager) ResetInfo(v *VerifyInfo) {
	tokenId := vm.GetId([]string{v.Username, v.Token}...)
	v.Number = 0
	vm.VerifyStore.Store(tokenId, v)
}

func (vm *VerifyManager) DelInfo(v *VerifyInfo) {
	tokenId := vm.GetId([]string{v.Username, v.Token}...)
	vm.VerifyStore.Delete(tokenId)
}
func (vm *VerifyManager) buildVerifyCode() string {
	code := make([]byte, vm.CodeLen)
	lettersLen := len(LETTERS)
	for i := 0; i < int(vm.CodeLen); i++ {
		code[i] = LETTERS[rand.Intn(lettersLen)]
	}
	return string(code)
}


type VerifyInstance struct {
	verifyInfo *VerifyInfo

	// 判断验证码
	Verify(verifyCode string) bool

	// 更新验证信息
	Update()

	// 重置验证信息
	Reset()
}


func (vi *VerifyInstance) UnAuth(userName string, token string) error {
	vi.verifyInfo = GetInfo(userName, token)
	vInfo = vInfo.(*VerifyInfo)
	if vInfo.Number < int32(LOCK_NUM) {
		return nil
	}
	interval := time.Now().Sub(codeInfo.CreateTime).Seconds()
	if interval < LOCK_USER_TIME.Seconds() {
		return errors.New("lock")
	}
	vInfo.Number = 0
	v.StoreData.Store(name, vInfo)
	return nil
}

func (v *VerifyManager) Verify(name string, code string) error {
	vInfo, _ := v.StoreData.Load(name)
	if vInfo.(*VerifyInfo).Number < LOGIN_NUM {
		return nil
	}

	if vInfo.Code == "" {
		vInfo.Code = GetCode(name)
		v.StoreData.Store(name, vInfo)
	}

	if code == vInfo.Code {
		return nil
	}

	return errors.New("verifyCode error")

}

func (v *VerifyManager) Update(name string) {
	vInfo, _ := v.StoreData.Load(name)
	vInfo.Number++
	if vInfo.Number == LOCK_NUM {
		vInfo.CreateTime = time.Now()
	}
	v.StoreData.Store(name)

}

func (v *VerifyManager) Reset(name string) {
	vInfo, _ := v.StoreData.Load(name)
	vInfo.Number = 0
	v.StoreData.Store(name)
}

func CreateVfCode() {

}
