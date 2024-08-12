package graph

type GraphNodePosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type GraphViewport struct {
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Zoom float32 `json:"zoom"`
}

type GraphNodeData struct {
	Text                string `json:"text"`
	Image               string `json:"image,omitempty"`
	HasFrameDataSection bool   `json:"hasFrameDataSection"`
}

type GraphNode struct {
	Data        GraphNodeData     `json:"data"`
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
	SourceX          float64     `json:"sourceX"`
	SourceY          float64     `json:"sourceY"`
	Target           string      `json:"target"`
	TargetX          float64     `json:"targetX"`
	TargetY          float64     `json:"targetY"`
	SourceHandle     string      `json:"sourceHandle"`
	TargetHandle     string      `json:"targetHandle"`
	InteractionWidth int         `json:"interactionWidth"`
}

type Graph struct {
	Nodes    []GraphNode   `json:"nodes"`
	Edges    []GraphEdge   `json:"edges"`
	Viewport GraphViewport `json:"viewport"`
}

// Returns a JSON marshaled graph. This graph is the starting point of all new files
func GetInitGraph() Graph {
	return Graph{
		Nodes: []GraphNode{
			{
				Id:          "1",
				Initialized: false,
				Position: GraphNodePosition{
					X: 25,
					Y: 90,
				},
				NodeType: "custom",
				Data: GraphNodeData{
					Text:                "Nouveau noeud",
					HasFrameDataSection: false,
				},
			},
		},
		Edges: []GraphEdge{},
		Viewport: GraphViewport{
			Zoom: 1,
			X:    0,
			Y:    0,
		},
	}

}
