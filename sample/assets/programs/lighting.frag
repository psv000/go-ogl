#version 410

uniform vec3 uLightPos;
uniform vec3 uLightColor;
uniform vec4 uColor;

in vec3 vFragPos;
in vec3 vNormal;
out vec4 frag_colour;

void main()
{
    float ambientStrength = 0.1;
    vec3 ambient = ambientStrength * uLightColor;

    vec3 norm = normalize(vNormal);
    vec3 lightDir = normalize(uLightPos - vFragPos);

    float diff = max(dot(norm, lightDir), 0.0);
    vec3 diffuse = diff * uLightColor;

    vec3 result = (ambient + diffuse) * uColor.rgb;
    frag_colour = vec4(result, uColor.a);
}