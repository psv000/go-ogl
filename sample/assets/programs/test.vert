#version 410
layout (location = 0) in vec3 vp;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

void main() {
    gl_Position = uProjection * uModel * uView * vec4(vp, 1.0);
}