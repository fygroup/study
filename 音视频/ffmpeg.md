### ffmpeg 命令
```
https://zhuanlan.zhihu.com/p/46903150
https://www.cnblogs.com/dwdxdy/p/3240167.html
```

### 视频流提取一帧
```
https://blog.csdn.net/leixiaohua1020/article/details/42181571?spm=1001.2014.3001.5501
https://blog.csdn.net/leixiaohua1020/article/details/39759623?spm=1001.2014.3001.5501
```

### time_base
```c++
// time_base的意思就是时间的刻度。如（1,25），那么时间刻度就是1/25；（1,9000），那么时间刻度就是1/90000

// AVStream
// AVStream->base_time单位为秒，通过avpriv_set_pts_info(st, 33, 1, 90000)函数，设置AVStream->time_base为1/90000。为什么是90000？因为mpeg的pts、dts都是以90kHz来采样的，所以采样间隔为1/90000秒。

// AVCodecContext
// AVCodecContext->time_base单位同样为秒，不过精度没有AVStream->time_base高，大小为1/framerate。


// 流转发的时间转换 demo
AVFormatContext* ifmt_ctx;
AVFormatContext* ofmt_ctx;
AVPacket* pkt;

AVRational in_video_time_base = ifmt_ctx->streams[video_index]->time_base;
AVRational in_video_frame_rate = ifmt_ctx->streams[video_index]->r_frame_rate;
int64_t calc_during = AV_TIME_BASE/(in_video_frame_rate.num * speed); // AV_TIME_BASE = 1s
int64 video_frame_index = 0;
while (1) {
    ret = av_read_frame(ifmt_ctx, pkt);
    pkt->pts = video_frame_index * calc_during) / (av_q2d(in_video_time_base) * AV_TIME_BASE);
    pkt->dts = pkt->pts;
    pkt->duration = calc_during / (av_q2d(in_video_time_base) * AV_TIME_BASE);
    video_frame_index++;

    AVStream* in_stream = ifmt_ctx->streams[pkt->stream_index];
    AVStream* out_stream = ofmt_ctx->streams[pkt->stream_index];
    av_packet_rescale_ts(pkt, in_stream->time_base, out_stream->time_base);

    av_interleaved_write_frame(ofmt_ctx, pkt);

}
```

### ffmpeg架构
```
```

### video audio 播放同步
```
```

### seek 原理
```
```

