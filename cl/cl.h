#define CL_USE_DEPRECATED_OPENCL_1_2_APIS
#if defined(__APPLE__)
#   include <OpenCL/cl.h>
#   include <OpenCL/cl_ext.h>
#   include <OpenCL/cl_gl.h>
#   include <OpenCL/cl_gl_ext.h>
#   include <OpenGL/CGLDevice.h>
#   include <OpenGL/CGLCurrent.h>
#   include <OpenCL/opencl.h>
#   include <OpenGL/OpenGL.h>
#   include <OpenGL/gl.h>
#else
#   include <CL/cl.h>
#   include <CL/cl_ext.h>
#   include <CL/cl_gl.h>
#endif
