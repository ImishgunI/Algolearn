<script setup lang="ts">
import { ref, watch, onMounted, computed } from "vue";
import * as d3 from "d3";

interface GraphData {
  nodes: string[];
  edges: { from: string; to: string }[];
}

interface GraphStep {
  graph: GraphData;
  activeNodes: string[];
  visitedNodes: string[];
}

const props = defineProps<{
  steps: GraphStep[];
  currentStep: number;
}>();

const svgRef = ref<SVGSVGElement>();
const current = computed(() => props.steps[props.currentStep]);

function draw() {
  if (!svgRef.value || !current.value) return;
  const svg = d3.select(svgRef.value);
  svg.selectAll("*").remove();
  const width = 400, height = 300;
  svg.attr("viewBox", `0 0 ${width} ${height}`);

  const { nodes, edges } = current.value.graph;
  const active = new Set(current.value.activeNodes);
  const visited = new Set(current.value.visitedNodes);

  // Симуляция расположения (фиксированная сетка для простоты)
  const nodePositions = new Map<string, { x: number; y: number }>();
  nodes.forEach((n, i) => {
    const angle = (i / nodes.length) * 2 * Math.PI;
    const r = Math.min(width, height) * 0.35;
    nodePositions.set(n, {
      x: width / 2 + r * Math.cos(angle),
      y: height / 2 + r * Math.sin(angle),
    });
  });

  // Рёбра
  svg.selectAll("line")
    .data(edges)
    .enter()
    .append("line")
    .attr("x1", d => nodePositions.get(d.from)!.x)
    .attr("y1", d => nodePositions.get(d.from)!.y)
    .attr("x2", d => nodePositions.get(d.to)!.x)
    .attr("y2", d => nodePositions.get(d.to)!.y)
    .attr("stroke", "#999")
    .attr("stroke-width", 2);

  // Вершины
  svg.selectAll("circle")
    .data(nodes)
    .enter()
    .append("circle")
    .attr("cx", n => nodePositions.get(n)!.x)
    .attr("cy", n => nodePositions.get(n)!.y)
    .attr("r", 20)
    .attr("fill", n => {
      if (active.has(n)) return "#f59e0b";
      if (visited.has(n)) return "#10b981";
      return "#6366f1";
    })
    .attr("stroke", "#fff")
    .attr("stroke-width", 2);

  // Подписи
  svg.selectAll("text")
    .data(nodes)
    .enter()
    .append("text")
    .attr("x", n => nodePositions.get(n)!.x)
    .attr("y", n => nodePositions.get(n)!.y + 4)
    .attr("text-anchor", "middle")
    .attr("fill", "#fff")
    .style("font-size", "12px")
    .text(n => n);
}

onMounted(draw);
watch([() => props.currentStep, () => props.steps], draw);
</script>

<template>
  <div class="graph-container">
    <svg ref="svgRef" width="100%" height="300"></svg>
  </div>
</template>

<style scoped>
.graph-container {
  background: rgba(0, 0, 0, 0.2);
  border-radius: 8px;
  padding: 8px;
}
</style>