import { Component, OnInit, ViewChild } from '@angular/core';
import { CgEdgeConfigService, OpcuaConfig, Node } from '../cg-edge-config.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent} from '../message-popup/message-popup.component';
import { saveAs } from "file-saver";

@Component({
  selector: 'app-opcua-mqtt-connector',
  templateUrl: './opcua-mqtt-connector.component.html',
  styleUrls: ['./opcua-mqtt-connector.component.css']
})
export class OpcuaMqttConnectorComponent implements OnInit {

  appName!: string;
  newTopic!: string;
  newNode: Node = new Node();
  opcuaConfig: OpcuaConfig = new OpcuaConfig();
  @ViewChild('file') file: any

  constructor(private CgEdgeConfigService: CgEdgeConfigService,
              public dialog: MatDialog) {}

  ngOnInit(): void {
    this.getConfig();
  }

  getConfig() {
    this.appName = "opcua-mqtt-connector"
    this.CgEdgeConfigService.getConfig(this.appName).subscribe((data) => {
      this.opcuaConfig = (data as OpcuaConfig)
    });
  }

  setConfig() {
    this.CgEdgeConfigService.setOpcuaConfig(this.opcuaConfig).subscribe((data) => {
      this.dialog.open(MessagePopupComponent, {data: {title: "Write Configuration", text: data}});
      this.getConfig()
    });
    
  }

  addNode() {
    this.newNode.Name = ""
    this.newNode.NodeID = ""
    this.opcuaConfig.NodesToRead.Nodes.push(this.newNode)
  }

  deleteNode() {
    this.opcuaConfig.NodesToRead.Nodes.splice(-1)
  }

  addPubTopic() {
    this.newTopic = "newtopic/sample"
    this.opcuaConfig.TopicsPub.Topic.push(this.newTopic)
  }

  deletePubTopic() {
    this.opcuaConfig.TopicsPub.Topic.splice(-1)
  }

  trackByFn(index: any, item: any) {
    return index;
 }

 importConfig() {
  this.file.nativeElement.click();
 }

 onFilesAdded() {
  const jsonfile = this.file.nativeElement.files[0];
  this.file.nativeElement.value = "";
  let fileReader  = new FileReader();
  fileReader.readAsText(jsonfile);
  fileReader.onload = () => {
    const jsonfiletext = fileReader.result
    let jsonObject: any = JSON.parse(jsonfiletext as string);
    let finalObject: OpcuaConfig = <OpcuaConfig>jsonObject;
    this.opcuaConfig = finalObject;
    this.setConfig();
  }
 }

 exportConfig() {
  let exportData = this.opcuaConfig;
  return saveAs(new Blob([JSON.stringify(exportData, null, 2)], { type: 'JSON' }), 'data.json');
}

}
