#ifndef _OPENCV3_XFEATURES2D_H_
#define _OPENCV3_XFEATURES2D_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/xfeatures2d.hpp>
extern "C" {
#endif

#include "$GOPATH/src/gocv.io/x/gocv/core.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::xfeatures2d::SIFT>* SIFT;
typedef cv::Ptr<cv::xfeatures2d::SURF>* SURF;
#else
typedef void* SIFT;
typedef void* SURF;
#endif

struct Mat SIFT_DrawMatches(SIFT d,
                            Mat img1,
                            KeyPoints &kp1,
                            Mat img2,
                            KeyPoints &kp2,
                            const std::vector<std::vector<DMatch>> &matches1to2,
                            InputOutputArray outImg,
                            const Scalar &matchColor = Scalar::all(-1),
                            const Scalar &singlePointColor = Scalar::all(-1),
                            const std::vector<std::vector<char>> &matchesMask = std::vector<std::vector<char>>(),
                            int flags = DrawMatchesFlags::DEFAULT
                            );

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_XFEATURES2D_H_