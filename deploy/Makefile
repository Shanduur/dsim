.ONESHELL:
.PHONY: test deps download build astyle

# OpenCV version to use.
OPENCV_VERSION?=4.5.0

# Temporary directory to put files into.
TMP_DIR?=/tmp/

# Build shared or static library
BUILD_SHARED_LIBS?=ON

# Package list for each well-known Linux distribution
RPMS=cmake curl wget git gtk2-devel libpng-devel libjpeg-devel libtiff-devel tbb tbb-devel libdc1394-devel unzip
DEBS=unzip wget build-essential cmake curl git libgtk2.0-dev pkg-config libavcodec-dev libavformat-dev libswscale-dev libtbb2 libtbb-dev libjpeg-dev libpng-dev libtiff-dev libdc1394-22-dev

explain:
	@echo "For quick install with typical defaults of OpenCV, run 'make install'"

# Detect Linux distribution
distro_deps=
ifneq ($(shell which dnf 2>/dev/null),)
	distro_deps=deps_fedora
else
ifneq ($(shell which apt-get 2>/dev/null),)
	distro_deps=deps_debian
else
ifneq ($(shell which yum 2>/dev/null),)
	distro_deps=deps_rh_centos
endif
endif
endif

# Install all necessary dependencies.
deps: $(distro_deps)

deps_rh_centos:
	sudo yum -y install pkgconfig $(RPMS)

deps_fedora:
	sudo dnf -y install pkgconf-pkg-config $(RPMS)

deps_debian:
	sudo apt-get -y update
	sudo apt-get -y install $(DEBS)


# Download OpenCV source tarballs.
download:
	rm -rf $(TMP_DIR)opencv
	mkdir $(TMP_DIR)opencv
	cd $(TMP_DIR)opencv
	curl -Lo opencv.zip https://github.com/opencv/opencv/archive/$(OPENCV_VERSION).zip
	unzip -q opencv.zip
	curl -Lo opencv_contrib.zip https://github.com/opencv/opencv_contrib/archive/$(OPENCV_VERSION).zip
	unzip -q opencv_contrib.zip
	rm opencv.zip opencv_contrib.zip
	cd -

# Build OpenCV.
build:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=OFF -D BUILD_opencv_java=NO -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D WITH_JASPER=OFF -DOPENCV_GENERATE_PKGCONFIG=ON ..
	$(MAKE) -j $(shell nproc --all)
	$(MAKE) preinstall
	cd -

# Build OpenCV on Raspbian with ARM hardware optimizations.
build_raspi:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=OFF -D BUILD_opencv_java=OFF -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D ENABLE_NEON=ON -D ENABLE_VFPV3=ON -D WITH_JASPER=OFF -D OPENCV_GENERATE_PKGCONFIG=ON ..
	$(MAKE) -j $(shell nproc --all)
	$(MAKE) preinstall
	cd -

# Build OpenCV on Raspberry pi zero which has ARMv6.
build_raspi_zero:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=OFF -D BUILD_opencv_java=OFF -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D ENABLE_VFPV2=ON -D WITH_JASPER=OFF -D OPENCV_GENERATE_PKGCONFIG=ON ..
	$(MAKE) -j $(shell nproc --all)
	$(MAKE) preinstall
	cd -

# Build OpenCV with non-free contrib modules.
build_nonfree:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=OFF -D BUILD_opencv_java=NO -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D WITH_JASPER=OFF -DOPENCV_GENERATE_PKGCONFIG=ON -DOPENCV_ENABLE_NONFREE=ON ..
	$(MAKE) -j $(shell nproc --all)
	$(MAKE) preinstall
	cd -

# Build OpenCV with openvino.
build_openvino:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D ENABLE_CXX11=ON -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D WITH_INF_ENGINE=ON -D InferenceEngine_DIR=/usr/local/dldt/inference-engine/build -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=OFF -D BUILD_opencv_java=NO -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D WITH_JASPER=OFF -DOPENCV_GENERATE_PKGCONFIG=ON -DOPENCV_ENABLE_NONFREE=ON ..
	$(MAKE) -j $(shell nproc --all)
	$(MAKE) preinstall
	cd -

# Build OpenCV with cuda.
build_all:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)
	mkdir build
	cd build
	rm -rf *
	cmake -j $(shell nproc --all) -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_SHARED_LIBS=${BUILD_SHARED_LIBS} -D ENABLE_CXX11=ON -D OPENCV_EXTRA_MODULES_PATH=$(TMP_DIR)opencv/opencv_contrib-$(OPENCV_VERSION)/modules -D WITH_INF_ENGINE=ON -D InferenceEngine_DIR=/usr/local/dldt/inference-engine/build -D BUILD_DOCS=OFF -D BUILD_EXAMPLES=OFF -D BUILD_TESTS=OFF -D BUILD_PERF_TESTS=OFF -D BUILD_opencv_java=NO -D BUILD_opencv_python=NO -D BUILD_opencv_python2=NO -D BUILD_opencv_python3=NO -D WITH_JASPER=OFF -DOPENCV_GENERATE_PKGCONFIG=ON -DWITH_CUDA=ON -DENABLE_FAST_MATH=1 -DCUDA_FAST_MATH=1 -DWITH_CUBLAS=1 -DCUDA_TOOLKIT_ROOT_DIR=/usr/local/cuda/ -DBUILD_opencv_cudacodec=OFF -D WITH_CUDNN=ON -D OPENCV_DNN_CUDA=ON -D CUDA_GENERATION=Auto ..
	$(MAKE) -j $(shell nproc --all)
	$(MAKE) preinstall
	cd -

# Cleanup old library files.
sudo_pre_install_clean:
	sudo rm -rf /usr/local/lib/cmake/opencv4/
	sudo rm -rf /usr/local/lib/libopencv*
	sudo rm -rf /usr/local/lib/pkgconfig/opencv*
	sudo rm -rf /usr/local/include/opencv*

# Do everything.
install: deps download sudo_pre_install_clean build sudo_install

# Do everything on Raspbian.
install_raspi: deps download build_raspi sudo_install

# Do everything on the raspberry pi zero.
install_raspi_zero: deps download build_raspi_zero sudo_install

# Install system wide.
sudo_install:
	cd $(TMP_DIR)opencv/opencv-$(OPENCV_VERSION)/build
	sudo $(MAKE) install
	sudo ldconfig
	cd -

astyle:
	astyle --project=.astylerc --recursive *.cpp,*.h