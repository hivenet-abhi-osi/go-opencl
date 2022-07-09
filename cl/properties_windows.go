package cl

// #include "cl.h"
// #define USE_GL_ATTACHMENTS (1)  // enable OpenGL attachments for Compute results
import "C"

type Property C.cl_context_properties

func GetInteropProperties(window GLFWHandle) []Property {
	return []Property{
		Property(C.CL_GL_CONTEXT_KHR), C.glfwGetWGLContext((*C.GLFWwindow)(window.Handle())),
		Property(C.CL_WGL_HDC_KHR), C.GetDC(C.glfwGetWin32Window((*C.GLFWwindow)(window.Handle()))),
		Property(C.CL_CONTEXT_PLATFORM), C.lPlatform(),
	}
}
