#include "drawing.h"

struct Mat SIFT_DrawMatches(SIFT d, Mat img1, KeyPoints &kp1, Mat img2, KeyPoints &kp2,
                            const std::vector<std::vector<DMatch>> &matches1to2, InputOutputArray outImg,
                            const Scalar &matchColor = Scalar::all(-1), const Scalar &singlePointColor = Scalar::all(-1),
                            const std::vector<std::vector<char>> &matchesMask = std::vector<std::vector<char>>(),
                            int flags = DrawMatchesFlags::DEFAULT) {
    Mat
}