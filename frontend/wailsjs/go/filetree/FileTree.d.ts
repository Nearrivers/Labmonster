// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {filetree} from '../models';
import {graph} from '../models';

export function CreateFile(arg1:string):Promise<filetree.Node>;

export function DeleteFile(arg1:string):Promise<void>;

export function DuplicateFile(arg1:string,arg2:string):Promise<string>;

export function GetDirectories():Promise<Array<string>>;

export function GetLabDirs():Promise<void>;

export function GetLabPath():Promise<string>;

export function GetRecentlyOpenedFiles():Promise<Array<string>>;

export function GetSubDirAndFiles(arg1:string):Promise<Array<filetree.Node>>;

export function MoveFileToExistingDir(arg1:string,arg2:string):Promise<string>;

export function OpenFile(arg1:string):Promise<graph.Graph>;

export function RenameFile(arg1:string,arg2:string,arg3:string):Promise<void>;

export function SaveFile(arg1:string,arg2:graph.Graph):Promise<void>;

export function SaveMedia(arg1:string,arg2:string,arg3:string):Promise<string>;
