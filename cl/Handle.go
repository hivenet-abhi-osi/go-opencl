package cl

import "unsafe"

//created to allow multiple GLFW instances play nice

type GLFWHandle interface {
	Handle() unsafe.Pointer
}
