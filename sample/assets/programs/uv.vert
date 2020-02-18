#version 400

layout (location = 0) in vec3 aPosition;
layout (location = 1) in vec2 aUVCoord;
layout (location = 2) in vec3 aNormal;

uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;

out vec4 vColor;

void main(void)
{
    gl_Position = uProjection * uModel * uView * vec4(aPosition, 1.);
}
