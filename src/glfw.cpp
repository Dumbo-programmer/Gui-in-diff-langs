#include <GLFW/glfw3.h>
.
. //Entry and glfwInit()
.
GLFWwindow* window = glfwCreateWindow(1000, 1000, "MyWindow", NULL, NULL);
glfwMakeContextCurrent(window);
