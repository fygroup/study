### 资料
```
http://qiushao.net/2020/05/17/IPCamera/himpp-overview/index.html
https://www.cnblogs.com/qinghaowusu/p/13610568.html
https://blog.csdn.net/weixin_40878354/article/details/106002321
https://zhuanlan.zhihu.com/p/187463660
https://github.com/licaibiao/hisi_sdk_develop
```

### 名词解释
```
DPU（Depth Process Unit）深度信息处理单元，主要是来实现双目测距和三维重建的硬件加速功能从而提高实时性
DMA 控制器（DMAC）直接在存储器和外设、外设和外设、存储器和存储器之间进行数据传输，避免CPU 干涉并减少了CPU 中断处理开销
LDC 镜头畸变校正
BAS(Bayer scaling) 即 Bayer 域缩放

// 编码
VEDU（Video Encode Decode Unit）是一个硬件实现的支持H.265/H.264 视频标准的编码器
// 解码
VDH(Video Decoding Module For High Definition)高清视频解码模块
PGD（PNG Decoder）是一个硬件加速模块

// 视频及图形处理
TDE（Two Dimensional Engine）利用硬件进行图形绘制，可以大大减少对CPU 的占用，同时提高了内存带宽的资源利用率
VPSS（Video Processing Sub System）视频处理子系统，实现视频处理功能。包含视频拼接、视频遮挡、视频马赛克处理、视频裁剪、缩放、亮度单分量处理、压缩、解压缩、mirror、flip 功能
VGS（Video Graphics System）视频图形系统，实现视频及图形处理功能。包含OSD 叠加、缩放、区域亮度和统计、视频裁剪、视频遮挡功能
GDC（Geometric Distortion Correction）几何畸变矫正，实现图像畸变矫正功能
AVSP（Any View Stitching Processing）全景拼接，实现最多六个镜头的拼接功能，包括720 度全景、360 度全景及非全景拼接都可实现
6-DOF-DIS（6-Degree Of Freedom-Digital Image Stabilization）六轴防抖，实现单目镜头下IPC、运动DV 和无人机等场景下的6 轴视频防抖功能
FRC 帧率控制,分为 2 种:组帧率控制和通道帧率控制
Crop 裁剪
NR（noise reduce） 去噪
cover 遮挡
scale 缩放
mirror 水平翻转
flip 垂直翻转
rotate 旋转
overlay 视频叠加
border 加边框
DEI 将交错隔行视频还原成逐行视频源
ES 边缘平滑
IE 图像增强
DCI 动态对比度


// 智能加速
IVE (Intelligent Video Engine)模块提供智能分析算法中所用到的一系列基础运算功能，以及部分耗时较大的特殊功能，是智能分析系统中的硬件加速模块
Vision DSP(Vision Processor 6)是给视觉处理加速的专用处理器，具有可编程的能力，基于Vision DSP 既可以开发供智能分析算法用的一系列基础运算功能，也可以实现复杂的算法
NNIE

// 视频接口
VI（Video Input），可以通过MIPI Rx（包含MIPI、LVDS、HiSPi、SLVS_EC）接口、BT.656/601、BT.1120 接口和DC（Digital Camera）接收视频数据，存入指定的内存区域。VI 内嵌ISP 图像处理单元，可以直接对接外部原始数据（BAYER RGB 数据）
VI 分成两个物理子模块：捕获子模块VICAP 和处理子模块VIPROC 组成
    VICAP 完成多路视频输入的数据捕获功能，并将捕获的数据存放到DDR 或者在线送给VIPROC
    VIPORC 用以支持离线模式（从DDR 读取数据）或者在线模式（从VICAP 接收在线）视频数据处理
SENSOR -> MIPI_RX -> VICAP -> VIROC0(ISP)
                           -> VIROC1(ISP)

VDP（Video Display Processor）模块主动从内存相应位置读取视频和图形数据，将视频层和图形层数据叠加后通过显示通道送出

// ISP
ISP 模块支持标准的Sensor 图像数据处理，包括自动白平衡、自动曝光、Demosaic、坏点矫正及镜头阴影矫正等基本功能，也支持WDR、DRC、降噪等高
级处理功能

VDA 视频侦测分析，通过检测视频的亮度变化，得出视频侦测分析结果。VDA 包含运动侦测（MD）和遮挡检测（OD）两种工作模式




```

### VI VPSS 工作模式
```
VI 和 VPSS 各自的工作模式分为在线 离线

// 在线模式	
    VI_CAP 与 VI_PROC 之间在线数据流传输,此模式下 VI_CAP不会写出 RAW 数据到 DDR,而是直接把数据流送给VI_PROC
    VI_PROC 与 VPSS 之间的在线数据流传输,在此模式下 VI_PROC不会写出 YUV 数据到 DDR,而是直接把数据流送给 VPSS
// 离线模式
    VI_CAP 写出 RAW 数据到DDR,然后 VI_PROC 从 DDR 读取 RAW 数据进行后处理
    VI_PROC 写出 YUV 数据到DDR,然后 VPSS 从 DDR 读取 YUV 数据进行后处理

在线模式时 VI 进行时序解析后直接在芯片内部将数据传递到VPSS，中间无DDR 写出的过程
在线模式可以省一定的带宽和内存，降低端到端的延时
但是，在线模式时，因为VI 不写出数据到DDR，无法进行CoverEx、OverlayEx、Rotate、LDC 等操作，需要在VPSS 各通道写出后再进行Rotate/LDC 等处理，而且有些功能只在离线下能支持，比如DIS
所以使用在线模式还是离线模式需要根据具体需求来决定。如果追求低延时，那自然要使用在线模式
```