#version 410
layout (location = 0) in vec3 vp;

uniform mat4 uModelViewProjection;

void main() {
    gl_Position = uModelViewProjection * vec4(vp, 1.0);
}