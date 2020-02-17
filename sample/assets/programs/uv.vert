#version 400

layout (location = 0) in vec3 aPosition;
layout (location = 1) in vec2 aUVCoord;
layout (location = 2) in vec3 aNormal;

uniform mat4 uModelViewProjection;

out vec4 vColor;

void main(void)
{
    gl_Position = uModelViewProjection * vec4(aPosition, 1.);
}
