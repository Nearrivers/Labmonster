export namespace filetree {
	
	export class Node {
	    name: string;
	    type: string;
	    // Go type: time
	    updatedAt: any;
	    extension: string;
	
	    static createFrom(source: any = {}) {
	        return new Node(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.type = source["type"];
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.extension = source["extension"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace graph {
	
	export class EdgeMarker {
	    color: string;
	    height: number;
	    width: number;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new EdgeMarker(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.color = source["color"];
	        this.height = source["height"];
	        this.width = source["width"];
	        this.type = source["type"];
	    }
	}
	export class GraphViewport {
	    x: number;
	    y: number;
	    zoom: number;
	
	    static createFrom(source: any = {}) {
	        return new GraphViewport(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.x = source["x"];
	        this.y = source["y"];
	        this.zoom = source["zoom"];
	    }
	}
	export class GraphEdge {
	    data: any;
	    id: string;
	    label: string;
	    markerEnd: EdgeMarker;
	    source: string;
	    sourceX: number;
	    sourceY: number;
	    target: string;
	    targetX: number;
	    targetY: number;
	    sourceHandle: string;
	    targetHandle: string;
	    interactionWidth: number;
	
	    static createFrom(source: any = {}) {
	        return new GraphEdge(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = source["data"];
	        this.id = source["id"];
	        this.label = source["label"];
	        this.markerEnd = this.convertValues(source["markerEnd"], EdgeMarker);
	        this.source = source["source"];
	        this.sourceX = source["sourceX"];
	        this.sourceY = source["sourceY"];
	        this.target = source["target"];
	        this.targetX = source["targetX"];
	        this.targetY = source["targetY"];
	        this.sourceHandle = source["sourceHandle"];
	        this.targetHandle = source["targetHandle"];
	        this.interactionWidth = source["interactionWidth"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GraphNodePosition {
	    x: number;
	    y: number;
	
	    static createFrom(source: any = {}) {
	        return new GraphNodePosition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.x = source["x"];
	        this.y = source["y"];
	    }
	}
	export class GraphNode {
	    data: any;
	    id: string;
	    initialized: boolean;
	    position: GraphNodePosition;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new GraphNode(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = source["data"];
	        this.id = source["id"];
	        this.initialized = source["initialized"];
	        this.position = this.convertValues(source["position"], GraphNodePosition);
	        this.type = source["type"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Graph {
	    nodes: GraphNode[];
	    edges: GraphEdge[];
	    viewport: GraphViewport;
	
	    static createFrom(source: any = {}) {
	        return new Graph(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nodes = this.convertValues(source["nodes"], GraphNode);
	        this.edges = this.convertValues(source["edges"], GraphEdge);
	        this.viewport = this.convertValues(source["viewport"], GraphViewport);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	

}

