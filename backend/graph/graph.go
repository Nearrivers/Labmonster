package graph

type GraphNodePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type GraphViewport struct {
	X    int     `json:"x"`
	Y    int     `json:"y"`
	Zoom float32 `json:"zoom"`
}

type GraphNode struct {
	Data        interface{}       `json:"data"`
	Id          string            `json:"id"`
	Initialized bool              `json:"initialized"`
	Position    GraphNodePosition `json:"position"`
	NodeType    string            `json:"type"`
}

type EdgeMarker struct {
	Color    string `json:"color"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
	EdgeType string `json:"type"`
}

type GraphEdge struct {
	Data             interface{} `json:"data"`
	Id               string      `json:"id"`
	Label            string      `json:"label"`
	MarkerEnd        EdgeMarker  `json:"markerEnd"`
	Source           string      `json:"source"`
	SourceX          int         `json:"sourceX"`
	SourceY          int         `json:"sourceY"`
	Target           string      `json:"target"`
	TargetX          int         `json:"targetX"`
	TargetY          int         `json:"targetY"`
	SourceHandle     string      `json:"sourceHandle"`
	TargetHandle     string      `json:"targetHandle"`
	InteractionWidth int         `json:"interactionWidth"`
}

type Graph struct {
	Nodes    []GraphNode   `json:"nodes"`
	Edges    []GraphEdge   `json:"edges"`
	Viewport GraphViewport `json:"viewport"`
}
