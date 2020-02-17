#version 410
layout (location = 0) in vec3 vp;
layout (location = 1) in vec4 vc;

uniform mat4 uModelViewProjection;

out vec4 color;

void main() {
    gl_Position = uModelViewProjection * vec4(vp, 1.0);
    color = vc;
}