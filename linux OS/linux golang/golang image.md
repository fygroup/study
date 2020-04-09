### 颜色
```
1、RGB
    三通道颜色
    通过红绿蓝三色通道, 外加alpha透明度, 来展示几乎所有的颜色
    RGB 255,255,255
    0为最暗, 255为最亮(白色)

2、HEX
    十六进制颜色
    通过16进制0~F这16个字符来表达颜色的
    000000为黑色, FFFFFF为白色

3、RGB转换为HEX
    RGB: 92,184,232
    92 / 16 = 5余12 -> 5C
    184 / 16 = 11余8 -> B8
    232 / 16 = 14余8 -> E8
    HEX: 5CB8E8

4、HEX转换RGB
    HEX: F26BC1
    F2 = 15和2 -> 15 * 16 + 2 = 242
    6B = 6和11 -> 6 * 16 + 11 = 107
    C1 = 12和1 -> 12 * 16 + 1 = 193
    RGB: 242,107,193

5、颜色渐变规律
    // 一次变换一个通道(当然可以多通道变换)
    红  255 0   0
            ↓
    黄  255 255 0
        ↓
    绿  0   255 0
                ↓
    青  0   255 255
            ↓
    蓝  0   0   255
        ↓
    粉  255 0   255

    // 最小到最大
    RGB 158,47,200
    47 < x < 200
```


### color
```
色彩    色彩模型 
Color   Model   

1、Color
    在计算机中一个color可以理解为一个像素三个通道RGB的叠加(A表示透明)
    (1) 接口
        // 方法返回预乘了alpha的红、绿、蓝色彩值和alpha通道值，范围都在[0, 0xFFFF]
        type Color interface {
            // 返回4个int32类型的值
            RGBA() (r, g, b, a uint32)
        }
    
    (2) 色彩
        1) Alpha
            type Alpha struct {
                A uint8
            }
            func (c Alpha) RGBA() (r, g, b, a uint32)
        2) Gray
            type Gray struct {
                Y uint8
            }
            func (c Gray) RGBA() (r, g, b, a uint32)
        3) RGBA
            type RGBA struct {
                R, G, B, A uint8
            }
            func (c RGBA) RGBA() (r, g, b, a uint32)

2、Model
    Model接口可以将任意Color接口转换为采用自身色彩模型的Color接口。转换可能会丢失色彩信息
    (1) 接口
        type Model interface {
            Convert(c Color) Color
        }
    (2) 如何定义Model
        func ModelFunc(f myFunc(Color) Color) Model
        // myFunc是自己定义的转换函数

        // 例如
        var YCbCrModel Model = ModelFunc(yCbCrModelMyFunc)
    (3) Palette
        // 代表一个色彩的调色板
        type Palette []Color
        // 返回调色板中与色彩c在欧几里德RGB色彩空间最接近的色彩
        func (p Palette) Convert(c Color) Color
        // 返回调色板中与色彩c在欧几里德RGB色彩空间最接近的色彩的索引
        func (p Palette) Index(c Color) int

```

### image
```
1、Image
    表示采用某色彩模型的颜色构成的有限矩形网格（即一幅图像）
    (1) 接口
        type Image interface {
            // ColorModel方法返回图像的色彩模型
            ColorModel() color.Model
            // Bounds方法返回图像的范围，范围不一定包括点(0, 0)
            Bounds() Rectangle
            // At方法返回(x, y)位置的色彩
            // At(Bounds().Min.X, Bounds().Min.Y)返回网格左上角像素的色彩
            // At(Bounds().Max.X-1, Bounds().Max.Y-1) 返回网格右下角像素的色彩
            At(x, y int) color.Color
        }


2、Point
    Point是X, Y坐标对。坐标轴是向右（X）向下（Y）的
    (1) 类型
        type Point struct {
            X, Y int
        }
        // 原点
        var ZP Point


3、Rectangle
    Rectangle代表一个矩形
    (1) 类型
        type Rectangle struct {
            Min, Max Point
        }

4、Image实例
    (1) Uniform
        Uniform类型代表一块面积无限大的具有同一色彩的图像
        type Uniform struct {
            C color.Color
        }
        func NewUniform(c color.Color) *Uniform

    (2) Alpha
        Alpha类型代表一幅内存中的图像
        type Alpha struct {
            // Pix保管图像的像素，内容为alpha通道值（即透明度）。
            // 像素(x, y)起始位置是Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1]
            Pix []uint8
            // Stride是Pix中每行像素占用的字节数
            Stride int
            // Rect是图像的范围
            Rect Rectangle
        }
        func NewAlpha(r Rectangle) *Alpha

    (3) Gray
        Gray类型代表一幅内存中的图像
    

image   PalettedImage   Config  Point   Rectangle




```