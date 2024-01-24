package cl

// #include "cl.h"
// #define USE_GL_ATTACHMENTS (1)  // enable OpenGL attachments for Compute results
// cl_context_properties getSG() {
//    CGLContextObj kCGLContext = CGLGetCurrentContext();
//    CGLShareGroupObj kCGLShareGroup = CGLGetShareGroup(kCGLContext);
//    return (cl_context_properties)kCGLShareGroup;
// }
import "C"

type Property C.cl_context_properties

func GetInteropProperties() []Property {
	kCGLShareGroup := C.getSG()
	return []Property{
		Property(C.CL_CONTEXT_PROPERTY_USE_CGL_SHAREGROUP_APPLE), Property(kCGLShareGroup),
	}
}
