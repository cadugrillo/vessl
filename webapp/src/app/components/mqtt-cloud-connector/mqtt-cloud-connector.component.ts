import { Component, OnInit, ViewChild } from '@angular/core';
import { VesslConfigService, MccConfig } from '../../services/vessl-config.service';
import {MatDialog} from '@angular/material/dialog';
import { MessagePopupComponent} from '../../popups/message-popup/message-popup.component';
import { saveAs } from "file-saver";

@Component({
  selector: 'app-mqtt-cloud-connector',
  templateUrl: './mqtt-cloud-connector.component.html',
  styleUrls: ['./mqtt-cloud-connector.component.css']
})
export class MqttCloudConnectorComponent implements OnInit {

  appName!: string;
  newTopic!: string;
  mccConfig: MccConfig = new MccConfig();
  @ViewChild('file') file: any

  constructor(private VesslConfigService: VesslConfigService,
              public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getConfig();
  }

  getConfig() {
    this.appName = "mqtt-cloud-connector"
    this.VesslConfigService.getConfig(this.appName).subscribe((data) => {
      this.mccConfig = (data as MccConfig);
    });
  }

  setConfig() {
    this.VesslConfigService.setMccConfig(this.mccConfig).subscribe((data) => {
      this.dialog.open(MessagePopupComponent, {data: {title: "Write Configuration", text: data}});
      this.getConfig()
    });
    
  }

  addSubTopic() {
    this.newTopic = "newtopic/sample"
    this.mccConfig.TopicsSub.Topic.push(this.newTopic);
  }

  deleteSubTopic() {
    this.mccConfig.TopicsSub.Topic.splice(-1)
  }

  addPubTopic() {
    this.newTopic = "newtopic/sample"
    this.mccConfig.TopicsPub.Topic.push(this.newTopic);
  }

  deletePubTopic() {
    this.mccConfig.TopicsPub.Topic.splice(-1);
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
    let finalObject: MccConfig = <MccConfig>jsonObject;
    this.mccConfig = finalObject;
    this.setConfig();
  }
 }

 exportConfig() {
  let exportData = this.mccConfig;
  return saveAs(new Blob([JSON.stringify(exportData, null, 2)], { type: 'JSON' }), 'data.json');
}

}

