// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {filetree} from '../models';
import {config} from '../models';
import {context} from '../models';

export function GetFileTree():Promise<Array<filetree.Node>>;

export function PrintTree():Promise<void>;

export function Same(arg1:filetree.Node):Promise<boolean>;

export function SetConfigFile(arg1:config.AppConfig):Promise<void>;

export function SetContext(arg1:context.Context):Promise<void>;

export function Walk(arg1:filetree.Node,arg2:any):Promise<void>;
