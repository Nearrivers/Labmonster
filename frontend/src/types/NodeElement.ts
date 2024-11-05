export interface NodeElement extends HTMLElement {
  dataset: {
    path: string;
    extension: string;
    type: 'file' | 'directory';
    file: string;
  };
}
