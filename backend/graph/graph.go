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

type GraphNodeData struct {
	Text                string `json:"text"`
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