export type Bar = {
  value: number;
  id: number;
};

export type SortStep = {
  type: "sort";

  bars: Bar[];

  active: number[];
  swapping: number[];

  lines: number[];

  pivot?: number[];
  partition?: number[];

  leftRange?: number[];
  rightRange?: number[];
  mergedRange?: number[];
};

export type GraphStep = {
  type: "graph";

  graph: {
    nodes: string[];
    edges: {
      from: string;
      to: string;
    }[];
  };

  activeNodes: string[];
  visitedNodes: string[];

  lines: number[];
};

export type Step = SortStep | GraphStep;