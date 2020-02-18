#version 410
layout (location = 0) in vec3 vp;
layout (location = 1) in vec4 vc;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

out vec4 color;

void main() {
    gl_Position = uProjection * uModel * uView * vec4(vp, 1.0);
    color = vc;
}