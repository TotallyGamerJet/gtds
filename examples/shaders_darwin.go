package main

const source = `#include <metal_stdlib>

using namespace metal;

struct Vertex {
	float4 position [[position]];
	float4 color;
};

vertex Vertex VertexShader(
	uint vertexID [[vertex_id]],
	device Vertex * vertices [[buffer(0)]],
	constant int2 * windowSize [[buffer(1)]],
	constant float2 * pos [[buffer(2)]]
) {
	Vertex out = vertices[vertexID];
	out.position.xy += *pos;
	float2 viewportSize = float2(*windowSize);
	out.position.xy = float2(-1 + out.position.x / (0.5 * viewportSize.x),
	                          1 - out.position.y / (0.5 * viewportSize.y));
	return out;
}

fragment float4 FragmentShader(Vertex in [[stage_in]]) {
	return in.color;
}
`
