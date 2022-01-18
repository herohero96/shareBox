<template>
  <div class="child-device" ref="childDevice">
    <div class="quick-set-config">
      <title-bar :titleName='"快速配置"'></title-bar>
      <div class="con">
        <p class="con-tip">导入配置后，并不可以立即生效，需要更改配置相关参数才会生效。</p>
        <div class="export-import">
          <div class="con-export">
            <g-button type="default">
              <a class='a' :href='`${baseUrl}system/exportSystemConfig`' download="config.json">导出配置</a>
            </g-button>
          </div>
          <div class="con-import">
            <div class="con-import-detail">
              <span class="name">{{currentFileName}}</span>
              <div class="file-upload">
                <g-button class="file-btn">导入配置</g-button>
                <input type="file" @change="fileChange($event)" class="file-input" accept=".json" />
              </div>
              <span class="error-tip" v-if="isImportFail"><i class="icon iconfont icon-warning"></i>导入配置失败</span>
            </div>            
          </div>
        </div>
      </div>
    </div>

    <div class="child-set-config">
      <title-bar :titleName='"设备配置"'></title-bar>
      <div class="con">
        <div class="device-list">
          <div
            v-for="(item, index) of portList"
            :key="index"
            @click="changeCurrentPort(item, index)"
            :class="{'device-item': true, 'selected': item.windowsName === currentPort.windowsName }"
          >
            <p class="name">
              <i v-if="item.importFix" class="icon iconfont icon-edit"></i>
              {{item.windowsName}}
            </p>
            <p class="cnt">{{item.cnt}}台</p>
          </div>
        </div>

        <div class="com-wrap" v-if="currentPort.windowsName && currentPort.windowsName.includes('COM')">
          <div class="com-con">
            <div class="com-name">
              <span>{{com.data.portName}}</span>
            </div>
            <div class="com-set">
              <div class="com-set-item">
                <label>波特率：</label>
                <Select v-model="com.data.baudrate" style="width:110px">
                  <Option v-for="(item, index) in com.baudrateList" :value="item" :key="index">{{ item }}</Option>
                </Select>
              </div>
              <div class="com-set-item">
                <label>数据位：</label>
                <Select v-model="com.data.dataBit" style="width:110px">
                  <Option v-for="(item, index) in com.dataBitList" :value="item" :key="index">{{ item }}</Option>
                </Select>
              </div>
              <div class="com-set-item">
                <label>停止位：</label>
                <Select v-model="com.data.stopBit" style="width:110px">
                  <Option v-for="(item, index) in com.stopBitList" :value="item" :key="index">{{ item }}</Option>
                </Select>
              </div>
              <div class="com-set-item">
                <label>奇偶校验位：</label>
                <Select v-model="com.data.parity" style="width:110px">
                  <Option v-for="item in com.parityList" :value="item.value" :key="item.key">{{ item.key }}</Option>
                </Select>
              </div>
              <div class="com-set-item">
                <g-button @click="modifyPortConfig">确定</g-button>
              </div>
            </div>
          </div>
        </div>

        <div style="height: 95px">
          <div class="option-wrap" :class="{'fix': isOptionHeaderFixed}" ref="optionHeader" >
            <div class="option-con">
              <div class="add">
                <g-button type="primary" @click="openAddModal">+ 新增设备</g-button>
              </div>
              <div class="delete" v-if="currentPort.device && currentPort.device.length">
                <Checkbox v-model="currentPort.selectAll" @on-change="checkAllDevice">全选</Checkbox>
                <Checkbox v-model="currentPort.selectInvert" @on-change="invertCheckDevice" style="margin-right: 34px;">反选</Checkbox>
                <g-button @click="openDeleteModal">删除设备</g-button>
              </div>
            </div>
          </div>
        </div>

        <div class="device-detail-wrap" ref="deviceDetailWrap">
          <div
            v-for="(item,index) of currentPort.device"
            :key="index"
            @click="changeDeviceItemSelect(item, index)"
            :class="{'device-detail-item': true, 'selected': item.isSelected}"
          >
            <div class="message">
              <span class="name">
                {{item.deviceName}}
                <i v-if="item.importFix" class="icon iconfont icon-edit"></i>
              </span>
              <span class="message-item">设备型号：{{item.type}}</span>
              <span class="message-item">
                <span :class="{'circle': true, 'green': !item.southError && !item.southConnectError, 'yellow' : ((!item.southError && item.southConnectError) || (item.southError && !item.southConnectError))}"></span>
                南向协议：{{item.southProtocol}}
              </span>
              <span class="message-item">
                <span :class="{'circle': true, 'green': !item.northError && !item.northConnectError, 'yellow' : ((!item.northError && item.northConnectError) || (item.northError && !item.northConnectError))}"></span>
                北向协议：{{item.northProtocol}}
              </span>
              <span class="message-item">接口：{{item.connectionPort}}</span>
            </div>
            <div class="nature">
              <!-- <span class="nature-item">氮气浓度：</span>
              <span class="nature-item">氮气浓度：</span> -->
            </div>
            <span class="option" @click="openEditModal(item)">编辑</span>
          </div>
        </div>

      </div>
    </div>

    <win-tip-second v-if="isImportSuccess" :content="'配置文件已经导入成功'"></win-tip-second>

    <win-common v-if="isShowDeleteModal" class="delete-win" :title="'确认删除以下所选设备？'">
      <div class="delete-win-con">
        <div class="delete-device-list">
          <span
            v-for="(item, index) of deleteDeviceList"
            :key="index"
            :class="{'delete-device-item': true, 'selected': item.isSelected}"
            @click="changeDeviceItemSelect(item, index)"
          >{{item.deviceName}}</span>
        </div>
        <div class="footer">
          <div class="cnt">共 {{deleteDeviceCnt}} 台设备</div>
          <div class="option">
            <g-button @click="cancelDeleteDevice" style="margin-right: 14px;">取消</g-button>
            <g-button @click="confirmDeleteDevice" type="primary">删除</g-button>
          </div>
        </div>
      </div>
    </win-common>

    <win-common v-if="isShowAddModal" class="add-win" :width="820" :isShowCloseIcon="true" @closeWin="confirmCloseAddModal">
      <div class="add-win-con" ref="addWinCon">
        <div class="child-device-info">
          <div class="title">
            <div class="name">设备信息</div>
            <div class="info">
              <Checkbox v-model="deviceInfo.data.template" label="保存为模板">保存为模板</Checkbox>
              <span class="remark">（保存为模板后，当前参数将作为模板自动填入下一个设备的配置页面）</span>
            </div>
          </div>
          <div class="content">
            <Form v-model="deviceInfo.data" ref="formChildDeviceInfo" :label-width="64">
              <div class="common">
                <FormItem label="设备名称：" style="display: inline-block;">
                  <Input
                    v-model="deviceInfo.data.deviceName"
                    @on-blur="checkDeviceInfoItem('deviceName')"
                    @on-change="deviceInfoDisabled = false"
                    :class="{'error': this.deviceInfo.check.deviceNameError}"
                    style="width: 160px;"
                  />
                  <p v-if="this.deviceInfo.check.deviceNameError" class="error-text">设备名称不能为空</p>
                </FormItem>
                <FormItem label="型号：" style="display: inline-block;">
                  <Input
                    v-model="deviceInfo.data.type"
                    @on-blur="checkDeviceInfoItem('type')"
                    @on-change="deviceInfoDisabled = false"
                    :class="{'error': this.deviceInfo.check.typeError}"
                    style="width: 160px;"
                  />
                  <p v-if="this.deviceInfo.check.typeError" class="error-text">型号不能为空</p>
                </FormItem>
                <FormItem label="位置：" style="display: inline-block;">
                  <Input
                    v-model="deviceInfo.data.installLocation"
                    @on-blur="checkDeviceInfoItem('installLocation')"
                    @on-change="deviceInfoDisabled = false"
                    :class="{'error': this.deviceInfo.check.installLocationError}"
                    style="width: 160px;"
                  />
                  <p v-if="this.deviceInfo.check.installLocationError" class="error-text">位置不能为空</p>
                </FormItem>
              </div>
              <div class="variable">
                <div class="inline-item">
                  <FormItem label="设备连接：">
                    <Select @on-change="deviceInfoDisabled = false" v-model="deviceInfo.data.connectionPort" style="width:120px">
                      <Option v-for="(item, index) of deviceInfo.connectionPortList" :value="item.windowsName" :key="index">{{item.windowsName}}</Option>
                    </Select>
                    <p v-if="this.deviceInfo.check.connectionPortError && !deviceInfo.data.connectionPort" class="error-text">请选择设备连接</p>
                  </FormItem>
                  <FormItem label="TCP：" :label-width="60" v-if="deviceInfo.data.connectionPort && !deviceInfo.data.connectionPort.includes('COM')">
                    <Select @on-change="deviceInfoDisabled = false" v-model="deviceInfo.data.tcpType" style="width:120px">
                      <Option v-for="(item, index) of deviceInfo.tcpTypeList" :value="item" :key="index">{{item}}</Option>
                    </Select>
                    <p v-if="this.deviceInfo.check.connectionPortError && !deviceInfo.data.tcpType" class="error-text">请选择 TCP</p>
                  </FormItem>
                </div>
                <div class="block-item" v-if="deviceInfo.data.connectionPort && !deviceInfo.data.connectionPort.includes('COM')">
                  <FormItem label="服务器IP：" v-if="deviceInfo.data.tcpType && deviceInfo.data.tcpType==='client'">
                    <Input
                      v-model="deviceInfo.data.sevrerIp"
                      @on-blur="checkDeviceInfoItem('sevrerIp')"
                      @on-change="deviceInfoDisabled = false"
                      :class="{'error': this.deviceInfo.check.sevrerIpError}"
                      style="width: 300px;"
                    />
                    <p v-if="this.deviceInfo.check.sevrerIpError" class="error-text">请输入正确的 IP</p>
                  </FormItem>
                  <FormItem label="port：" v-if="deviceInfo.data.tcpType">
                    <Input
                      v-model="deviceInfo.data.serverPort"
                      @on-blur="checkDeviceInfoItem('serverPort')"
                      @on-change="deviceInfoDisabled = false"
                      :class="{'error': this.deviceInfo.check.serverPortError}"
                      style="width: 300px;"
                    />
                    <p v-if="this.deviceInfo.check.serverPortError" class="error-text">请输入正确的端口号</p>
                  </FormItem>
                </div>
              </div>
            </Form>
          </div>
          <div class="footer">
            <g-button
              @click="saveDeviceInfo"
              :disabled="!deviceInfo.data.deviceName || (deviceInfo.data.uuid !== '' && deviceInfoDisabled)"
              :loading="deviceInfoLoading"
            >保存</g-button>
          </div>
        </div>
        <div class="south">
          <div class="title">
            <div class="name">南向协议</div>
            <!-- <div class="info">
              <Checkbox @on-change="changeSaveTemplate" v-model="deviceInfo.data.southTemplate" label="保存为模板">保存为模板</Checkbox>
              <span class="remark">（保存为模板后，当前参数将作为模板自动填入下一个设备的配置页面）</span>
            </div> -->
          </div>
          <div class="content">
            <Form v-model="southInfo" ref="formSouthInfo" :label-width="64">
              <div class="treaty-wrap">
                <FormItem label="协议选择：">
                  <Select @on-change="southInfoDisabled = false" v-model="deviceInfo.data.southProtocol" style="width:160px">
                    <Option v-for="(item, index) of southInfo.southProtocolList" :value="item.value" :key="index">{{item.key}}</Option>
                  </Select>
                </FormItem>
              </div>
              <!-- 南向平田 -->
              <Form v-model="southInfo.hirataData" v-if="deviceInfo.data.southProtocol === 'HIRATA'" ref="formSouthInfoHirataData" :label-width="64">
                <div class="time-wrap">
                  <FormItem label="扫描时间：">
                    <Input
                      v-model="southInfo.hirataData.scanTime"
                      @on-blur="checkSouthInfoHirataDataItem('scanTime')"
                      @on-change="southInfoDisabled = false"
                      :class="{'error': this.southInfo.hirataData.scanTimeError}"
                      style="width: 80px;"
                      /> s
                      <p v-if="this.southInfo.hirataData.scanTimeError" class="error-text">扫描时间为大于等于 1 的整数</p>
                  </FormItem>
                </div>
              </Form>
              <!-- 南向 MQTT -->
              <Form v-model="southInfo.modbusData" v-if="deviceInfo.data.southProtocol === 'MQTT'" ref="formSouthInfoModbusData" :label-width="64">
                <div class="time-wrap">
                  <FormItem label="扫描时间：">
                    <Input v-model="southInfo.modbusData.time" style="width: 80px;" /> ms
                  </FormItem>
                </div>
                <div class="variable">
                  <div class="inline-item">
                    <FormItem label="设备ID：" :label-width="62">
                      <Input v-model="southInfo.modbusData.childDeviceId" style="width: 160px;" />
                    </FormItem>
                    <FormItem label="起始地址：" :label-width="62">
                      <Input v-model="southInfo.modbusData.startAddress" style="width: 160px;" />
                    </FormItem>
                    <FormItem label="功能码：" :label-width="52">
                      <Select v-model="southInfo.modbusData.funcCode" style="width:70px">
                        <Option v-for="(item, index) of southInfo.funcCodeList" :value="item" :key="index">{{item}}</Option>
                      </Select>
                    </FormItem>
                    <FormItem label="寄存器个数：" :label-width="76">
                      <Input v-model="southInfo.modbusData.registerCnt" style="width: 60px;" />
                    </FormItem>
                  </div>
                  <div v-for="(item,index) of southInfo.modbusData.variableList" :key="index" class="variable-item">
                    <div class="row">
                      <FormItem label="变量名称：">
                        <Input v-model="item.name" style="width: 160px;" />
                      </FormItem>
                      <FormItem label="变量类型：">
                        <Select v-model="item.type" style="width:160px">
                          <Option v-for="(type, index) of item.typeList" :value="type" :key="index">{{type}}</Option>
                        </Select>
                      </FormItem>
                      <FormItem class="btn" :label-width="0">
                        <i @click="addSouthVariable" class="icon iconfont icon-plus"></i>
                        <Poptip
                          placement="top-end"
                          class="delete-poptip-set"
                          confirm
                          offset="15"
                          @on-ok="deleteSouthVariable(item, index)"
                          @on-cancel="cancelDeleSouthVariable"
                          ok-text="Yes"
                          cancel-text="No"
                        >
                          <i class="icon iconfont icon-minus"></i>
                          <div slot="title">
                            <i class="iconfont icon-warn-fill"></i>
                            <span>确认要删除吗？</span>
                          </div>
                      </Poptip>
                      </FormItem>
                    </div>
                    <div class="row">
                      <FormItem label="变量单位：">
                        <Input v-model="item.unit" style="width: 160px;" />
                      </FormItem>
                      <FormItem label="变量处理：">
                        <Select v-model="item.handle" style="width:160px">
                          <Option v-for="handle of item.handleList" :value="handle.value" :key="handle.key">{{handle.key}}</Option>
                        </Select>
                      </FormItem>
                      <FormItem :label-width="0" v-if="item.handle==='jingdu'">
                        <Select v-model="item.jingdu" style="width:160px">
                          <Option v-for="(jingdu, index) of item.jingduList" :value="jingdu" :key="index">{{jingdu}}</Option>
                        </Select>
                      </FormItem>
                      <FormItem :label-width="0" v-if="item.handle==='chushu'">
                        <Input v-model="item.chushu" style="width: 160px;" />
                      </FormItem>
                    </div>
                  </div>
                </div>
              </Form>
              <!-- 南向三协 -->
              <Form v-model="southInfo.Sankyo" v-if="deviceInfo.data.southProtocol === 'SANKYO'" ref="formSouthInfoSankyoData" :label-width="140">
           
                <div class="time-wrap">
                  <FormItem label="Select Buffering Item：">
                    <Select v-model="southInfo.sankyoData.bufferingItem" 
                    @on-change="southInfoDisabled = false"
                    style="width:160px">
                      <Option v-for="(item, index) of southInfo.sankyoData.bufferingList" :value="item" :key="index">{{item}}</Option>
                    </Select>
                  </FormItem>
                </div>

                <div class="time-wrap">
                  <FormItem label="Interval：">
                    <Input
                      v-model="southInfo.sankyoData.collectIntervalTime"
                      @on-blur="checkSouthInfoSankyoDataItem('collectIntervalTime')"
                      @on-change="southInfoDisabled = false"
                      :class="{'error': this.southInfo.sankyoData.collectIntervalTimeError}"
                      style="width: 80px;"
                      /> 
                      <p v-if="this.southInfo.sankyoData.collectIntervalTimeError" class="error-text">采集时间间隔为4或者8</p>
                  </FormItem>
                </div> 
                <div class="time-wrap">
                  <FormItem label="Bufferin：">
                    <Input
                      v-model="southInfo.sankyoData.collectNum"
                      @on-blur="checkSouthInfoSankyoDataItem('collectNum')"
                      @on-change="southInfoDisabled = false"
                      :class="{'error': this.southInfo.sankyoData.collectNumError}"
                      style="width: 80px;"
                      />
                      <p v-if="this.southInfo.sankyoData.collectNumError" class="error-text">采集个数为大于等于4且小于等于2000的整数</p>
                  </FormItem>
                </div>
              </Form>

            </Form>
          </div>
          <div class="footer">
            <g-button
              @click="saveSouthInfo"
              :disabled="deviceInfo.data.uuid === '' || (deviceInfo.data.uuid && southInfoDisabled)"
              :loading="southInfoLoading"
            >保存</g-button>
          </div>
        </div>
        <div class="north">
          <div class="title">
            <div class="name">北向协议</div>
            <!-- <div class="info">
              <Checkbox @on-change="changeSaveTemplate" v-model="deviceInfo.data.northTemplate" label="保存为模板">保存为模板</Checkbox>
              <span class="remark">（保存为模板后，当前参数将作为模板自动填入下一个设备的配置页面）</span>
            </div> -->
          </div>
          <div class="content">
            <Form v-model="northInfo" ref="formNorthInfo" :label-width="64">
              <div class="mqtt-wrap" v-if="northInfo.mqttDataList.length > 0">
                <div class="mqtt-item" v-for="(item, index) of northInfo.mqttDataList" :key="index" >
                  <div class="option">
                    <div class="type">
                      <span style="margin-right: 4px; vertical-align: sub;" :class="{'circle': true, 'green': item.testMqtt}"></span>
                      <span>MQTT</span>
                    </div>
                    <div class="btn">
                      <span class="test-mqtt" @click="testMqtt(item)">测试</span>
                      <Poptip
                        placement="top-end"
                        class="delete-poptip-set"
                        confirm
                        offset="15"
                        @on-ok="deleteMQTTNorthboundProtocol(item, index)"
                        ok-text="Yes"
                        cancel-text="No"
                      >
                        <i class="icon iconfont icon-minus"></i>
                        <div slot="title">
                          <i class="iconfont icon-warn-fill"></i>
                          <span>确认要删除吗？</span>
                        </div>
                    </Poptip>
                    </div>
                  </div>
                  <div class="info">
                    <div class="row">
                      <FormItem label="服务器Ip：">
                        <Input
                          v-model="item.mqttServerIp"
                          @on-blur="checkNorthInfoMqttItem(item, 'mqttServerIp', index)"
                          @on-change="northInfoDisabled = false"
                          :class="{'error': item.mqttServerIpError}"
                          style="width: 160px;"
                        />
                        <p v-if="item.mqttServerIpError" class="error-text">请输入正确的 IP</p>
                      </FormItem>
                      <FormItem label="Port：">
                        <Input
                          v-model="item.mqttServerPort"
                          @on-blur="checkNorthInfoMqttItem(item, 'mqttServerPort', index)"
                          @on-change="northInfoDisabled = false"
                          :class="{'error': item.mqttServerPortError}"
                          style="width: 160px;"
                        />
                        <p v-if="item.mqttServerPortError" class="error-text">请输入正确的端口</p>
                      </FormItem>
                    </div>
                    <div class="row row-bottom">
                      <FormItem label="ClientID：">
                        <Input
                          v-model="item.mqttServerClientID"
                          @on-blur="checkNorthInfoMqttItem(item, 'mqttServerClientID', index)"
                          @on-change="northInfoDisabled = false"
                          :class="{'error': item.mqttServerClientIDError}"
                          style="width: 160px;"
                        />
                        <p v-if="item.mqttServerClientIDError" class="error-text">请输入正确的 ClientID</p>
                      </FormItem>
                      <FormItem label="UserName：">
                        <Input
                          v-model="item.username"
                          @on-change="northInfoDisabled = false"
                          :class="{'error': item.usernameError}"
                          style="width: 160px;"
                        />
                      </FormItem>
                      <FormItem label="Password：">
                        <Input
                          v-model="item.password"
                          @on-change="northInfoDisabled = false"
                          :class="{'error': item.passwordError}"
                          style="width: 160px;"
                        />
                      </FormItem>
                    </div>
                    <!-- 当前位置 -->
                    <div class="topic-list">
                      <div v-for="(topic, topicIndex) of item.mqttServerTopic" :key="topicIndex" class="row">
                        <FormItem
                          :label-width="deviceInfo.data.southProtocol === 'HIRATA' ? 70 : deviceInfo.data.southProtocol === 'SANKYO' ? 85 : 64"
                          :label="(deviceInfo && deviceInfo.data.southProtocol !== 'HIRATA' && deviceInfo.data.southProtocol !== 'SANKYO') ? 'Topic：' : ''"
                          style="position: relative"
                        >
                          <span :class="{'topic-has-axle-wrap' : true, hirata: deviceInfo.data.southProtocol === 'HIRATA'}">
                            <span class="topic-has-axle-container">
                              <span class="topic-has-axle" v-if="(deviceInfo && deviceInfo.data.southProtocol === 'HIRATA') || (deviceInfo && deviceInfo.data.southProtocol === 'SANKYO')">Topic:</span>
                              <span class="topic-axle" v-if="deviceInfo && deviceInfo.data.southProtocol === 'SANKYO'">{{deviceInfo.data.southProtocol === 'SANKYO' &&  northInfo.topicSankyoDescribe[topicIndex] ? northInfo.topicSankyoDescribe[topicIndex] : `轴${topicIndex + 1}`}}</span>
                              <span class="topic-axle" v-if="deviceInfo && deviceInfo.data.southProtocol === 'HIRATA'">{{deviceInfo.data.southProtocol === 'HIRATA' &&  northInfo.topicHirataDescribe[topicIndex] ? northInfo.topicHirataDescribe[topicIndex] : `J${topicIndex + 1}`}}</span>
                            </span>
                          </span>
                          <Input
                            v-model="item.mqttServerTopic[topicIndex]"
                            @on-blur="checkNorthInfoMqttItem(item, 'topic', index)"
                            @on-change="northInfoDisabled = false"
                            :class="{'error': item.mqttServerTopicError[topicIndex]}"
                            style="width: 407px;"
                          />
                          <p v-if="item.mqttServerTopicError[topicIndex]" class="error-text">Topic 不能为空</p>
                        </FormItem>
                        <div class="topic-btn">
                          <i v-if="topicIndex == item.mqttServerTopic.length - 1" class="icon iconfont icon-plus" @click="item.mqttServerTopic.push(''),northInfoDisabled = false"></i>
                          <Poptip
                            v-if="topicIndex == item.mqttServerTopic.length - 1 && topicIndex > 0"
                            placement="top-end"
                            class="delete-poptip-set"
                            confirm
                            offset="15"
                            @on-ok="item.mqttServerTopic.splice(topicIndex, 1),northInfoDisabled = false"
                            ok-text="Yes"
                            cancel-text="No"
                          >
                            <i class="icon iconfont icon-minus"></i>
                            <div slot="title">
                              <i class="iconfont icon-warn-fill"></i>
                              <span>确认要删除吗？</span>
                            </div>
                          </Poptip>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="tibco-wrap" v-if="northInfo.tibcoRvDataList.length > 0">
                <div class="tibco-item" v-for="(item, index) of northInfo.tibcoRvDataList" :key="index">
                  <div class="option">
                    <div class="type">
                      <!-- <span style="margin-right: 4px; vertical-align: sub;" :class="{'circle': true, 'open': !deviceInfo.data.northError}"></span> -->
                      <span>TibcoRV</span>
                    </div>
                    <div class="btn">
                      <Poptip
                        placement="top-end"
                        class="delete-poptip-set"
                        confirm
                        offset="15"
                        @on-ok="deleteTibcoRvNorthboundProtocol(item, index)"
                        @on-cancel="cancelDeleteTibcoRvNorthboundProtocol"
                        ok-text="Yes"
                        cancel-text="No"
                      >
                        <i class="icon iconfont icon-minus"></i>
                        <div slot="title">
                          <i class="iconfont icon-warn-fill"></i>
                          <span>确认要删除吗？</span>
                        </div>
                    </Poptip>
                    </div>
                  </div>
                  <div class="info">
                    <div class="row">
                      <FormItem label="服务：">
                        <Input
                          v-model="item.service"
                          @on-blur="checkNorthInfoTibcoItem(item, 'service', index)"
                          @on-change="northInfoDisabled = false"
                          :class="{'error': item.serviceError}"
                          style="width: 160px;"
                        />
                        <p v-if="item.serviceError" class="error-text">服务不能为空</p>
                      </FormItem>
                      <FormItem label="网络：">
                        <Input
                          v-model="item.network"
                          @on-blur="checkNorthInfoTibcoItem(item, 'network', index)"
                          @on-change="northInfoDisabled = false"
                          :class="{'error': item.networkError}"
                          style="width: 160px;"
                        />
                        <p v-if="item.networkError" class="error-text">网络不能为空</p>
                      </FormItem>
                      <FormItem label="守护信息：">
                        <Input
                          v-model="item.daemon"
                          @on-blur="checkNorthInfoTibcoItem(item, 'daemon', index)"
                          @on-change="northInfoDisabled = false"
                          :class="{'error': item.daemonError}"
                          style="width: 160px;"
                        />
                        <p v-if="item.daemonError" class="error-text">守护信息不能为空</p>
                      </FormItem>
                      <FormItem label="主题：">
                        <Input
                          v-model="item.subject"
                          @on-blur="checkNorthInfoTibcoItem(item, 'subject', index)"
                          @on-change="northInfoDisabled = false"
                          :class="{'error': item.subjectError}"
                          style="width: 160px;"
                        />
                        <p v-if="item.subjectError" class="error-text">主题不能为空</p>
                      </FormItem>
                    </div>
                  </div>
                </div>
              </div>
            </Form>
          </div>
          <div class="north-footer">
            <!-- <Poptip placement="right-end" width="84">
              <div slot="content">
                <p @click="addNorthMqtt" class="treaty">MQTT</p>
                <p @click="addNorthTibco" class="treaty">TibcoRV</p>
              </div>
              <g-button type="primary">+ 新增</g-button>
            </Poptip> -->
            <div class="poptip">
              <!-- <g-button type="primary" @click.stop="toggleAddNorthPoptip">+ 新增</g-button> -->
              <Button type="primary" @click.stop="toggleAddNorthPoptip">新增</Button>
              <div class="tip" v-show="northInfo.isShowAddNorthPoptip" ref="poptipContent">
                <span class="triangle"></span>
                <p @click="addNorthMqtt">MQTT</p>
                <p @click="addNorthTibco">TibcoRV</p>
              </div>
            </div>
            <g-button
              @click="saveNorthInfo"
              :disabled="!deviceInfo.data.uuid || (northInfo.tibcoRvDataList.length <= 0 && northInfo.mqttDataList.length <= 0) || (this.deviceInfo.data.uuid && northInfoDisabled)"
              :loading="northInfoLoading"
            >保存</g-button>
          </div>
        </div>
        <div>
          <g-button :disabled="!deviceInfo.data.uuid" type="primary" @click="beginComm">确定</g-button>
        </div>
      </div>
    </win-common>

    <win-common :isShowMark="true" class='close-confirm-win' v-if="isShowCloseConfirmWin" :title="'关闭确认'">
      <div class="close-confirm-con">
        <p class='close-tip'>当前参数还未保存，是否离开？</p>
        <div class='close-btn clearfix'>
          <g-button class="btn-cancel fr" type="primary" @click="isShowCloseConfirmWin = false">继续编辑</g-button>
          <g-button class="fr" type="default" @click="closeAddModal">离开</g-button>
        </div>
      </div>
    </win-common>

  </div>
</template>

<script lang="ts">
import { Vue, Component, Watch } from 'vue-property-decorator';
import config from '@/config';
import GButton from '@/components/button/g-button.vue';
import TitleBar from '@/components/title-bar/title-bar.vue';
import WinTipSecond from '@/components/windows/win-tip-second.vue';
import WinCommon from '@/components/windows/win-common.vue';
import {
  importSystemConfigApi,
  exportSystemConfigApi,
  queryPortInfoApi,
  queryPortConfigApi,
  modifyPortConfigApi,
  queryAllDeviceApi,
  addDeviceApi,
  queryDeviceApi,
  editDeviceApi,
  deleteDeviceApi,
  addMQTTNorthboundProtocolApi,
  queryMQTTNorthboundProtocolApi,
  editMQTTNorthboundProtocolApi,
  deleteMQTTNorthboundProtocolApi,
  addTibcoRvNorthboundProtocolApi,
  editTibcoRvNorthboundProtocolApi,
  deleteTibcoRvNorthboundProtocolApi,
  queryTibcoRvNorthboundProtocolApi,
  testMqttConnectApi,
  beginCommApi,
  addHirataProtocolApi,
  editHirataProtocolApi,
  queryHirataProtocolApi,
  addSankyoProtocolApi,
  querySankyoProtocolApi,
  editSankyoProtocolApi,
  deleteSankyoProtocolApi,
  queryCommStateApi,
} from '@/api/child-device/child-device';

@Component({
  components: {
    GButton,
    TitleBar,
    WinTipSecond,
    WinCommon,
  },
})
export default class ChildDevice extends Vue {
  private deviceInfoDisabled: boolean = true; // 设备信息保存按钮 disabled
  private deviceInfoLoading: boolean = false; // 设备信息保存按钮 loading
  private southInfoDisabled: boolean = true; // 南向保存按钮 disabled
  private southInfoLoading: boolean = false; // 南向保存按钮 loading
  private northInfoDisabled: boolean = true; // 北向保存按钮 disabled
  private northInfoLoading: boolean = false; // 北向保存按钮 loading
  private baseUrl: string = config.baseUrl;
  private currentFileName: string = ''; // 当前配置 file name
  private isImportFail: boolean = false; // 导入系统配置失败
  private isImportSuccess: boolean = false; // 导入系统配置成功

  private com: any = { // com
    baudrateList: [110, 300, 600, 1200, 2400, 4800, 9600, 14400, 19200, 38400, 56000, 115200], // 波特率
    dataBitList: [4, 5, 6, 7, 8], // 数据位
    stopBitList: [1, 1.5, 2], // 停止位
    parityList: [ // 奇偶校验
      {
        key: 'NONE',
        value: 0,
      },
      {
        key: 'ODD',
        value: 1,
      },
      {
        key: 'EVEN',
        value: 2,
      },
      {
        key: 'MARK',
        value: 3,
      },
      {
        key: 'SPACE',
        value: 4,
      },
    ],
    data: {
      portName: '',
      baudrate: '',
      dataBit: '',
      stopBit: '',
      parity: '',
    },
  };

  private portList: any = []; // 端口列表
  private currentPortIndex: any = 0; // 当前端口下标
  private currentPort: any = {}; // 当前选中的端口
  private allDevice: any = []; // 所有设备列表
  private queryCommStateTimer: any; // 查询相关设备通讯状态定时器

  private optionOffsetTop: any; // 新增、删除设备按钮滚动高度
  private optionOffsetHeight: any; // 新增、删除设备按钮高度
  private mainContentCon: any; // 父级元素
  private isOptionHeaderFixed: boolean = false;
  private scrollDiv: any; // 滚动事件的div
  private isRunning: boolean = false; // 节流标识


  private isShowDeleteModal: boolean = false; // 删除设备弹框
  private deleteDeviceList: any = []; // 删除设备列表
  private deleteDeviceCnt: number = 0; // 删除设备个数

  private isShowAddModal: boolean = false; // 新增设备弹框
  private deviceInfo: any = { // 新增设备——设备信息
    data: {
      connectionPort: '', // 设备的连接端口
      createTime: '', // 创建时间
      deviceName: '', // 设备的名称
      importFix: false,
      installLocation: '', // 设备的安装位置
      mqttnorthboundList: [], // MQTT北向协议的id列表
      serverPort: '', // 设备连接的服务器端口
      sevrerIp: '', // 设备连接的serverIp
      southProtocol: '', // 南向协议的枚举类型
      southboundList: [], // 南向协议的协议id列表
      tcpType: '', // 设备连接的TCP种类
      template: true, // 设备信息是否存为模板
      northTemplate: true, // 北向是否存为模板
      southTemplate: true, // 南向是否存为模板
      tibocoRvnorthboundList: [], // TibcoRv北向协议的id列表
      type: '', // 设备的型号
      uuid: '', // 设备的唯一识别码
    },
    tcpTypeList: ['server', 'client'], // 设备连接的TCP种类列表
    connectionPortList: [], // 设备的连接端口列表
    check: {
      deviceNameError: false,
      typeError: false,
      installLocationError: false,
      connectionPortError: false,
      tcpTypeError: false,
      serverPortError: false,
      sevrerIpError: false,
    },
  };
  private deviceInfoOrigin = JSON.parse(JSON.stringify(this.deviceInfo));
  private deviceInfoDataOrigin = JSON.parse(JSON.stringify(this.deviceInfo.data)); // 设备信息参数副本

  private northInfo: any = { // 新增设备——北向协议
    isShowAddNorthPoptip: false,
    template: true, // 是否存为模板
    mqttDataList: [], // MQTT北向协议的列表
    checkNorthInfoMqttObj: {mqttServerIpError: false, mqttServerPortError: false, mqttServerClientIDError: false, passwordError: false, usernameError: false},
    mqttData: {
      type: 'MQTT', // 协议类型
      deviceId: '', // MQTT北向通讯协议所绑定的设备唯一识别码
      mqttServerClientID: '', // MQTT的ClientID
      mqttServerIp: '', // MQTT的服务器IP
      mqttServerPort: '', // MQTT的服务器的端口
      mqttServerToken: '', // MQTT服务器的Token
      mqttServerTopic: [''], // MQTT服务器的Topic
      testMqtt: false, // 测试 MQTT
      password: '', // MQTT服务器的密码
      username: '', // MQTT服务器的用户名
      uuid: '',
      mqttServerIpError: false,
      mqttServerPortError: false,
      mqttServerClientIDError: false,
      mqttServerTopicError: [],
      passwordError: false,
      usernameError: false,
    },
    tibcoRvDataList: [],
    checkNorthInfoTibcoObj: {serviceError: false, networkError: false, daemonError: false, subjectError: false},
    tibcoRvData: {
      type: 'TibcoRV',
      deviceId: '', // Tibco Rv 该北向通讯协议所绑定的设备唯一识别码
      uuid: '', // Tibco Rv 北向通讯协议的唯一识别码
      service: '', // 服务
      network: '', // 网络
      daemon: '', // 守护信息
      subject: '', // 主题
      serviceError: false,
      networkError: false,
      daemonError: false,
      subjectError: false,
    },
    topicHirataDescribe: ['J1-X轴', 'J2-Y轴', 'J3-Z1', 'J4-C', 'J5', 'J6-Z2', 'J7-L1', 'J8-L2'], // 平田轴号描述信息
    topicSankyoDescribe: ['轴1-上手臂', '轴2-下手臂', '轴3-Z2', '轴4-Z1', '轴5-TH', '轴6-X'], // 三协轴号描述信息
  };
  private northInfoOrigin = JSON.parse(JSON.stringify(this.northInfo));
  private northInfoMqttDataOrigin = JSON.parse(JSON.stringify(this.northInfo.mqttData)); // 北向协议 MQTT 参数副本
  private northInfoTibcoDataOrigin = JSON.parse(JSON.stringify(this.northInfo.tibcoRvData)); // 北向协议 TibcoRv 参数副本

  private southInfo: any = { // 新增设备——南向协议
    southProtocolList: [ // 协议类型列表s
      {
        key: '透传',
        value: 'TRANSPARENT',
      },
      {
        key: '平田协议',
        value: 'HIRATA',
      },
      {
        key: '三协协议',
        value: 'SANKYO',
      },
    ],
    template: true, // 是否存为模板
    funcCodeList: ['01', '02', '03', '04'],
    handleList: [
      {
        key: '精度',
        value: 'jingdu',
      },
      {
        key: '除数',
        value: 'chushu',
      },
      {
        key: '开根号',
        value: 'genhao',
      },
    ],
    variableEmpty: {
      name: '',
      type: '',
      typeList: ['int', 'float'],
      unit: '',
      jingdu: '',
      jingduList: ['1', '0.1', '0.01', '0.001'],
      handle: 'jingdu',
      handleList: [
        {
          key: '精度',
          value: 'jingdu',
        },
        {
          key: '除数',
          value: 'chushu',
        },
        {
          key: '开根号',
          value: 'genhao',
        },
      ],
      chushu: '',
    },
    modbusData: {
      time: '',
      childDeviceId: '',
      startAddress: '',
      funcCode: '',
      handle: 'jingdu',
      registerCnt: '',
      variableList: [
        {
          name: '',
          type: '',
          typeList: ['int', 'float'],
          unit: '',
          jingdu: '',
          jingduList: ['1', '0.1', '0.01', '0.001'],
          handle: 'jingdu',
          handleList: [
            {
              key: '精度',
              value: 'jingdu',
            },
            {
              key: '除数',
              value: 'chushu',
            },
            {
              key: '开根号',
              value: 'genhao',
            },
          ],
          chushu: '',
        },
      ],
    },
    hirataData: {
      scanTime: '',
      uuid: '',
      scanTimeError: false,
    },
    sankyoData: {
      bufferingList: ['1:Speed1&PosErr&TrqCmd', '2:Speed2&State&TrqCmd', '3:SpeedF&TrqCmd', '4:Speed3&RMS&Thermal', '5:Speed4&PosErr&TrqExt', '6:Speed5&StateA&StateB'],
      bufferingItem: '1:Speed1&PosErr&TrqCmd',
      collectIntervalTime: '8',
      collectIntervalTimeError: false,
      collectNum: '2000',
      collectNumError: false,
    },
  };
  private southInfoOrigin = JSON.parse(JSON.stringify(this.southInfo));
  private southInfoHirataDataOrigin = JSON.parse(JSON.stringify(this.southInfo.hirataData));

  private isShowCloseConfirmWin: boolean = false;

  private mounted() {
    // 监听滚动事件
    this.handleScroll();
    // 查询端口信息
    this.queryPortInfo();
  }

  private beforeDestroy() {
    this.hide();
  }

  private destroyed() {
    this.scrollDiv.removeEventListener('scroll', this.scrollFn);
    clearInterval(this.queryCommStateTimer);
  }

  // 导入配置文件
  private fileChange(event: any) {
    const file = event.target.files[0];
    if (file) {
      const fileName = file.name;
      const formData = new FormData();
      formData.append('systemConfig', file);
      importSystemConfigApi(formData).then((res: any) => {
        if (res.status === 200) {
          this.currentFileName = fileName;
          console.log(this.currentFileName);
          this.isImportSuccess = true;
          const timer = setTimeout(() => {
            this.isImportSuccess = false;
            this.queryPortInfo();
          }, 3000);
        } else {
          this.isImportFail = true;
          const timer = setTimeout(() => {
            this.isImportFail = false;
          }, 3000);
        }
      });
    }
  }

  // 排序
  private compare(a: any, b: any) {
    if (a.createTime > b.createTime) {
      return -1;
    } else if (a.createTime < b.createTime) {
      return 1;
    } else {
      return 0;
    }
  }
  // 对象转数组，按时间倒序排序
  private objToArr(data: any) {
    const arr: any = [];
    Object.keys(data).forEach((d: any) => {
      arr.push(data[d]);
    });

    return arr.sort(this.compare);
  }

  // 查询端口信息,和端口对应的设备信息
  private queryPortInfo() {
    queryPortInfoApi().then((res: any) => {
      const tmpPortList = res.data;
      tmpPortList.unshift({windowsName: '全部'});
      queryAllDeviceApi().then((device: any) => {
        this.allDevice = this.objToArr(device.data);
        tmpPortList.forEach((port: any) => {
          port.device = [];
          port.cnt = 0;
          port.importFix = false;
          this.allDevice.forEach((d: any) => {
            d.northProtocol = '';
            d.isSelected = false;
            d.selectAll = false;
            d.selectInvert = false;
            if (d.mqttnorthboundList && d.mqttnorthboundList.length > 0) {
              d.northProtocol += 'MQTT';
            }
            if (d.tibocoRvnorthboundList && d.tibocoRvnorthboundList.length > 0) {
              d.northProtocol += ',TibcoRv';
            }
            if (d.connectionPort === port.windowsName || port.windowsName === '全部') {
              port.cnt ++;
              if (port.windowsName !== '全部' && d.importFix) {
                port.importFix = true;
              }
              port.device.push(d);
            }
          });
        });
        this.portList = tmpPortList;
        this.deviceInfo.connectionPortList = JSON.parse(JSON.stringify(this.portList));
        this.deviceInfo.connectionPortList.shift();
        this.currentPort = JSON.parse(JSON.stringify(this.portList[this.currentPortIndex]));
        clearInterval(this.queryCommStateTimer);
        this.queryCommState();
        this.mutifyQueryCommState();
      });
    });
  }

  // 查询设备通讯状态
  private queryCommState() {
    if (this.currentPort && this.currentPort.device.length > 0) {
      this.currentPort.device.forEach((d: any, index: any) => {
        queryCommStateApi(d.uuid).then((res: any) => {
          this.$set(this.currentPort.device[index], 'northError', res.data.northError);
          this.$set(this.currentPort.device[index], 'northConnectError', res.data.northConnectError);
          this.$set(this.currentPort.device[index], 'southError', res.data.southError);
          this.$set(this.currentPort.device[index], 'southConnectError', res.data.southConnectError);
        });
      });
    }
  }

  private mutifyQueryCommState() {
    if (this.currentPort && this.currentPort.device.length > 0) {
      this.queryCommStateTimer = setInterval(() => {
        this.queryCommState();
      }, 3000);
    }
  }

  // 查询端口配置
  private queryPortConfig(portName: string) {
    queryPortConfigApi(portName).then((res: any) => {
      if (res.status === 200) {
        this.com.data = JSON.parse(JSON.stringify(res.data));
      }
    });
  }

  // 修改端口配置
  private modifyPortConfig() {
    modifyPortConfigApi(this.com.data).then((res: any) => {
      if (res.status === 200) {
        this.$Message.info('修改成功');
      }
    });
  }

  // change 当前选中的端口
  private changeCurrentPort(item: any, index: any) {
    this.currentPort = JSON.parse(JSON.stringify(item));
    this.currentPortIndex = index;
    if (item.windowsName.includes('COM')) {
      this.queryPortConfig(item.windowsName);
    }
    clearInterval(this.queryCommStateTimer);
    this.queryCommState();
    this.mutifyQueryCommState();
  }

  // 滚动事件
  private handleScroll() {
    this.scrollDiv = document.getElementsByClassName('main-content-con')[0];
    this.scrollDiv.addEventListener('scroll', this.scrollFn);
  }
  private scrollFn() {
    if (this.isRunning) {
      return;
    }
    this.isRunning = true;
    window.requestAnimationFrame( () => {
      if (!this.isOptionHeaderFixed) {
        this.optionOffsetTop = (this as any).$refs.optionHeader.offsetTop;
      }
      this.$nextTick( () => {
        const layoutHeaderHeight: number = (document as any).getElementsByClassName('ivu-layout-header')[0].offsetHeight;
        this.isOptionHeaderFixed = this.scrollDiv.scrollTop > (this.optionOffsetTop - layoutHeaderHeight) ? true : false;
      } );
      this.isRunning = false;
    } );
  }

  // 全选
  private checkAllDevice(event: any) {
    this.currentPort.selectInvert = false;
    this.currentPort.device.forEach((item: any) => {
      item.isSelected = event;
    });
  }

  // 反选
  private invertCheckDevice(event: any) {
    this.currentPort.selectAll = false;
    this.currentPort.device.forEach((item: any) => {
      item.isSelected = !item.isSelected;
    });
  }

  // 更新设备选中状态
  private changeDeviceItemSelect(item: any, index: any) {
    item.isSelected = !item.isSelected;
    item.isSelected ? this.deleteDeviceCnt ++ : this.deleteDeviceCnt --;
    this.$set(this.currentPort.device[index], 'selectInvert', false);
    const selectAll = this.currentPort.device.every( ({ isSelected }: any) => isSelected );
    selectAll ? this.$set(this.currentPort, 'selectAll', true) : this.$set(this.currentPort, 'selectAll', false);
  }

  // 打开删除设备弹框
  private openDeleteModal() {
    this.deleteDeviceList = this.currentPort.device.filter( ( {isSelected}: any) => isSelected);
    this.deleteDeviceCnt = this.deleteDeviceList.length;
    if (this.deleteDeviceCnt > 0) {
      this.isShowDeleteModal = true;
    } else {
      this.$Message.error(`请选择设备`);
    }
  }

  // 确认删除设备
  private async confirmDeleteDevice() {
    const result: any = [];
    this.deleteDeviceList.forEach((item: any) => {
      if (item.isSelected) {
        result.push({uuid: item.uuid, deviceName: item.deviceName});
      }
    });

    // 删除
    const totalCnt = result.length;
    const resultList: any = [];
    for (let i = 0; i < totalCnt; i++) {
      resultList.push(await deleteDeviceApi(result[i].uuid));
    }

    // 处理结果
    this.isShowDeleteModal = false;
    this.queryPortInfo();
    const errorArray: any = [];
    resultList.forEach((res: any, index: any) => {
      if (res.status !== 200) {
        errorArray.push(result[index].deviceName);
      }
    });
    if (errorArray.length === 0 && totalCnt > 0) {
      this.$Message.info('删除成功');
    } else {
      this.$Message.error(`${errorArray.toString()} 删除失败`);
    }
  }

  // 取消删除设备
  private cancelDeleteDevice() {
    this.isShowDeleteModal = false;
  }

  // 打开新增设备弹框
  private openAddModal() {
    this.isShowAddModal = true;
    clearInterval(this.queryCommStateTimer); // 停止查询列表灯的状态
    const templateDevice = this.portList[0] && this.portList[0].device.length > 0 && this.portList[0].device[0] || '';
    this.saveInfoToLocal();
    if (templateDevice) { // 获取模板
      if (templateDevice.template) {
        this.deviceInfoDisabled = false;
        this.southInfoDisabled = false;
        this.northInfoDisabled = false;
        this.queryDevice(templateDevice, true); // 查询设备
        if (templateDevice.southProtocol === 'HIRATA') {
          this.queryHirataProtocol(templateDevice, true); // 查询南向平田
        }
        this.queryMQTTNorthboundProtocol(templateDevice, true); // 查询 MQTT 北向协议
        this.queryTibcoRvNorthboundProtocol(templateDevice, true); // 查询 TibcoRV 北向协议
      }
    }
  }

  // 打开编辑设备弹框
  private openEditModal(item: any) {
    this.isShowAddModal = true;
    clearInterval(this.queryCommStateTimer); // 停止查询列表灯的状态
    this.queryDevice(item , false); // 查询设备
    if (item.southProtocol === 'HIRATA') {
      this.queryHirataProtocol(item, false); // 查询南向平田
    }
    if (item.southProtocol === 'SANKYO') {
      console.log('item', item);

      this.querySankyoProtocol(item);

    }

    this.queryMQTTNorthboundProtocol(item, false); // 查询 MQTT 北向协议
    this.queryTibcoRvNorthboundProtocol(item, false); // 查询 TibcoRV 北向协议
  }

  // 保存参数到 sessionStorage
  private saveInfoToLocal() {
    setTimeout(() => {
      sessionStorage.setItem('deviceInfoData', JSON.stringify(this.deviceInfo.data));
      if (this.deviceInfo.data.southProtocol === 'HIRATA') { // 平田
        sessionStorage.setItem('southInfoHirataData', JSON.stringify(this.southInfo.hirataData));
      } else if (this.deviceInfo.data.southProtocol === 'SANKYO') { // 三协
        //
      }
      sessionStorage.setItem('mqttDataList', JSON.stringify(this.northInfo.mqttDataList));
      sessionStorage.setItem('tibcoRvDataList', JSON.stringify(this.northInfo.tibcoRvDataList));
    }, 100);
  }

  // 判断是否关闭新增设备弹框
  private confirmCloseAddModal() {
    const oldDeviceInfoData = sessionStorage.getItem('deviceInfoData');
    const oldSouthInfoHirataData = sessionStorage.getItem('southInfoHirataData');
    const oldMqttDataList = sessionStorage.getItem('mqttDataList');
    const oldTibcoRvDataList = sessionStorage.getItem('tibcoRvDataList');

    if (oldDeviceInfoData !== JSON.stringify(this.deviceInfo.data) ||
      oldMqttDataList !== JSON.stringify(this.northInfo.mqttDataList) ||
      oldTibcoRvDataList !== JSON.stringify(this.northInfo.tibcoRvDataList) ||
      (this.deviceInfo.data.southProtocol === 'HIRATA' && oldSouthInfoHirataData !== JSON.stringify(this.southInfo.hirataData))) {
        console.log(this.southInfo.hirataData);
        this.isShowCloseConfirmWin = true;
    } else {
      this.closeAddModal();
    }
  }

  // 关闭新增设备弹框
  private closeAddModal() {
    this.isShowCloseConfirmWin = false;
    this.isShowAddModal = false;
    this.queryPortInfo();
    this.deviceInfoDisabled = true;
    this.southInfoDisabled = true;
    this.northInfoDisabled = true;
    this.deviceInfo = JSON.parse(JSON.stringify(this.deviceInfoOrigin));
    this.northInfo = JSON.parse(JSON.stringify(this.northInfoOrigin));
    this.southInfo = JSON.parse(JSON.stringify(this.southInfoOrigin));
    sessionStorage.removeItem('deviceInfoData');
    sessionStorage.removeItem('southInfoHirataData');
    sessionStorage.removeItem('mqttDataList');
    sessionStorage.removeItem('tibcoRvDataList');
  }

  // 确定新增设备弹窗
  private beginComm() {
    if (this.deviceInfo.data.uuid) {
      beginCommApi(this.deviceInfo.data.uuid).then((res: any) => {
        if (res.status === 200) {
          this.$Message.info('通讯成功');
          this.closeAddModal();
        }
      });
    }
  }

  // 校验设备信息
  private checkDeviceInfoItem(type: any) {
    let isValid: boolean = true;
    const reg = /^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/;
    if (type === 'sevrerIp') {
      if (this.deviceInfo.data.sevrerIp && reg.test(this.deviceInfo.data.sevrerIp)) {
        this.deviceInfo.check[`${type}Error`] = false;
      } else {
        this.deviceInfo.check[`${type}Error`] = true;
        isValid = false;
      }
    } else if (type === 'serverPort') {
      if (this.deviceInfo.data.serverPort > 0 && this.deviceInfo.data.serverPort <= 65535) {
        this.deviceInfo.check[`${type}Error`] = false;
      } else {
        this.deviceInfo.check[`${type}Error`] = true;
        isValid = false;
      }
    } else {
      if (this.deviceInfo.data[type]) {
        this.deviceInfo.check[`${type}Error`] = false;
      } else {
        this.deviceInfo.check[`${type}Error`] = true;
        isValid = false;
      }
    }

    return isValid;
  }
  // 校验设备信息全部参数
  private checkDeviceInfoAll() {
    let isValid: boolean = true;
    const checkDeviceInfoList = ['deviceName', 'type', 'installLocation', 'connectionPort'];
    if (this.deviceInfo.data.connectionPort) {
      if (!this.deviceInfo.data.connectionPort.includes('COM')) {
        checkDeviceInfoList.push('tcpType');
        if (this.deviceInfo.data.tcpType) {
          checkDeviceInfoList.push('serverPort');
          if (this.deviceInfo.data.tcpType === 'client') {
            checkDeviceInfoList.push('sevrerIp');
          }
        }
      }
    }
    checkDeviceInfoList.forEach((item: any) => {
      if (!this.checkDeviceInfoItem(item)) {
        isValid = false;
      }
    });

    return isValid;
  }

  // 保存设备信息
  private saveDeviceInfo() {
    if (!this.deviceInfoLoading && this.checkDeviceInfoAll()) {
      this.deviceInfoLoading = true;
      if (this.deviceInfo.data.uuid) { // 编辑
        editDeviceApi({...this.deviceInfo.data, importFix: false}).then((res: any) => {
          if (res.status === 200) {
            this.deviceInfo.data = {...res.data};
            this.$Message.info('保存成功');
            this.saveInfoToLocal();
            this.deviceInfoDisabled = true;
            this.deviceInfoLoading = false;
          }
        });
      } else { // 新增
        addDeviceApi(this.deviceInfo.data).then((res: any) => {
          if (res.status === 200) {
            this.deviceInfo.data = {...res.data};
            this.$Message.info('保存成功');
            this.saveInfoToLocal();
            this.deviceInfoDisabled = true;
            this.deviceInfoLoading = false;
          }
        });
      }
    } else {
      this.$Message.error('参数错误');
    }
  }

  // 查询设备信息
  private queryDevice(device: any, isTemplate: boolean) {
    queryDeviceApi(device.uuid).then((res: any) => {
      if (isTemplate) {
        this.deviceInfo.data = {...res.data, uuid: '', hirataId: '', mqttnorthboundList: [], tibocoRvnorthboundList: []};
        this.saveInfoToLocal();
      } else {
        this.deviceInfo.data = {...res.data};
        this.checkDeviceInfoAll(); // 校验参数
        this.saveInfoToLocal();
      }
      console.log(this.deviceInfo.data);
    });
  }

  // 校验南向平田
  private checkSouthInfoHirataDataItem(type: any) {
    let isValid: boolean = true;
    if (type === 'scanTime') {
      if (!isNaN(this.southInfo.hirataData.scanTime) && Math.floor(this.southInfo.hirataData.scanTime) === Number(this.southInfo.hirataData.scanTime) && this.southInfo.hirataData.scanTime >= 1) {
        this.southInfo.hirataData[`${type}Error`] = false;
        isValid = true;
      } else {
        this.southInfo.hirataData[`${type}Error`] = true;
        isValid = false;
      }
    }

    return isValid;
  }
  private checkSouthInfoHirataDataAll() {
    let isValid: boolean = true;
    const checkSouthInfoHirataDataList = ['scanTime'];
    checkSouthInfoHirataDataList.forEach((item: any) => {
      if (!this.checkSouthInfoHirataDataItem(item)) {
        isValid = false;
      }
    });

    return isValid;
  }

  // 获取南向平田参数
  private getSouthHirataParam() {
    const param = {
      scanTime: this.southInfo.hirataData.scanTime,
      deviceId: this.deviceInfo.data.uuid,
      uuid: this.southInfo.hirataData.uuid ? this.southInfo.hirataData.uuid : '',
    };

    return param;
  }

  // 校验南向三协
  private checkSouthInfoSankyoDataItem(type: string) {
    if (type === 'collectIntervalTime') {
      this.southInfo.sankyoData.collectIntervalTimeError = ! ( [4, 8].includes(this.southInfo.sankyoData.collectIntervalTime * 1) );
    }
    if (type === 'collectNum') {
      const num = this.southInfo.sankyoData.collectNum;
      this.southInfo.sankyoData.collectNumError = ! (num >= 4 && num <= 2000 && /^[1-9]\d*$/.test(num));
    }

  }
  private checkSouthInfoSankyoDataAll() {
    const errorList = ['collectIntervalTime', 'collectNum'];
    errorList.forEach( (item) => {
      this.checkSouthInfoSankyoDataItem(item);
    } );
    const isValid: boolean = errorList.every((item) => {
      return !this.southInfo.sankyoData[`${item}Error`];
    });

    return isValid;

  }

  // 获取南向三协参数
  private getSouthSankyoParam() {
    const param = {
      bufferingItem: this.southInfo.sankyoData.bufferingItem,
      collectIntervalTime: this.southInfo.sankyoData.collectIntervalTime * 1,
      collectNum: this.southInfo.sankyoData.collectNum * 1,
      deviceId: this.deviceInfo.data.uuid,
      uuid: this.deviceInfo.data.sankyoId ? this.deviceInfo.data.sankyoId : '',
    };

    return param;
  }


  // 编辑设备
  private editDeviceInfo() {
    editDeviceApi(this.deviceInfo.data).then((device: any) => {
      if (device.status === 200) {
        this.deviceInfo.data = {...device.data};
        this.$Message.info('保存成功');
        this.saveInfoToLocal();
        this.southInfoLoading = false;
      }
    });
  }

  // 保存南向
  private saveSouthInfo() {
    if (this.deviceInfo.data.southProtocol === 'HIRATA') {
      // 平田
      if (!this.southInfoLoading && this.checkSouthInfoHirataDataAll()) {
        this.southInfoLoading = true;
        if (this.southInfo.hirataData.uuid) {
          editHirataProtocolApi(this.getSouthHirataParam()).then((res: any) => {
            this.southInfo.hirataData.uuid = res.data.uuid;
            this.editDeviceInfo();
            this.southInfoDisabled = true;
          });
        } else {
          addHirataProtocolApi(this.getSouthHirataParam()).then((res: any) => {
            this.southInfo.hirataData.uuid = res.data.uuid;
            this.deviceInfo.data.hirataId = res.data.uuid;
            this.editDeviceInfo();
            this.southInfoDisabled = true;
          });
        }
      }
    } else if (this.deviceInfo.data.southProtocol === 'SANKYO') {
      // 三协
      if (!this.southInfoLoading && this.checkSouthInfoSankyoDataAll()) {
        this.southInfoLoading = true;
        if (this.deviceInfo.data.sankyoId) {
          editSankyoProtocolApi(this.getSouthSankyoParam()).then((res: any) => {
            this.deviceInfo.data.sankyoId = res.data.uuid;
            this.editDeviceInfo();
            this.southInfoDisabled = true;
          });
        } else {
          addSankyoProtocolApi(this.getSouthSankyoParam()).then((res: any) => {
            this.deviceInfo.data.sankyoId = res.data.uuid;
            this.editDeviceInfo();
            this.southInfoDisabled = true;
          });
        }
      }
    } else if (this.deviceInfo.data.southProtocol === 'TRANSPARENT') {
      // 透传
      if (!this.southInfoLoading) {
        this.southInfoLoading = true;
        this.editDeviceInfo();
        this.southInfoDisabled = true;
      }
    } else if (!this.deviceInfo.data.southProtocol) {
      this.$Message.error('请选择协议');
    }
  }

  // 查询南向平田
  private queryHirataProtocol(device: any, isTemplate: boolean) {
    if (device.hirataId) {
      queryHirataProtocolApi(device.hirataId).then((res: any) => {
        this.southInfo.hirataData.scanTime = res.data.scanTime;
        if (!isTemplate) {
          this.southInfo.hirataData.uuid = res.data.uuid;
        }
        this.checkSouthInfoHirataDataAll();
        this.saveInfoToLocal();
      });
    }
  }

  // 查询三协
  private querySankyoProtocol(device: any) {
    console.log(device);
  
    if (device) {
      querySankyoProtocolApi(device.sankyoId).then( (res: any) => {
        console.log(res);
        this.southInfo.sankyoData.bufferingItem = res.data.bufferingItem
        this.southInfo.sankyoData.collectIntervalTime = res.data.collectIntervalTime
        this.southInfo.sankyoData.collectNum = res.data.collectNum
      } )
    }

  }


  // 新增南向 MQTT
  private addSouthVariable() {
    this.southInfo.modbusData.variableList.push(JSON.parse(JSON.stringify(this.southInfo.variableEmpty)));
  }

  // 删除南向 MQTT
  private deleteSouthVariable(param: any, index: any) {
    //
  }

  // 取消删除南向 MQTT
  private cancelDeleSouthVariable() {
    //
  }

  // 获取 MQTT 北向协议参数
  private getNorthMqttParam(mqttItem: any) {
    const param = {
      deviceId: this.deviceInfo.data.uuid, // MQTT北向通讯协议所绑定的设备唯一识别码
      mqttServerIp: mqttItem.mqttServerIp, // MQTT的服务器IP
      mqttServerPort: mqttItem.mqttServerPort, // MQTT的服务器的端口
      mqttServerClientID: mqttItem.mqttServerClientID, // MQTT的ClientID
      password: mqttItem.password, // MQTT服务器的密码
      username: mqttItem.username, // MQTT服务器的用户名
      mqttServerTopic: mqttItem.mqttServerTopic, // MQTT服务器的Tpopic
      uuid: mqttItem.uuid,
      template: this.northInfo.template,
    };

    return param;
  }

  // 获取 TibcoRV 北向协议参数
  private getNorthTibcoRVParam(tibcoItem: any) {
    const param = {
      deviceId: this.deviceInfo.data.uuid,
      uuid: tibcoItem.uuid,
      service: tibcoItem.service,
      network: tibcoItem.network,
      daemon: tibcoItem.daemon,
      subject: tibcoItem.subject,
    };

    return param;
  }

  // 添加/编辑 MQTT/TibcoRV 北向协议
  private saveNorthInfo() {
    const mqttTotalCnt = this.northInfo.mqttDataList.length;
    const tibcoTotalCnt = this.northInfo.tibcoRvDataList.length;
    let mqttSuccessCnt = 0;
    let tibcoSuccessCnt = 0;
    const tipType = mqttTotalCnt >= tibcoTotalCnt ? 'mqtt' : 'tibcorv';

    // mqtt
    if (!this.northInfoLoading) {
      this.northInfo.mqttDataList.forEach((mqttItem: any, index: any) => {
        if (this.checkNorthInfoMqttAll(mqttItem, index)) {
          this.northInfoLoading = true;
          if (mqttItem.uuid) {
            editMQTTNorthboundProtocolApi(this.getNorthMqttParam(mqttItem)).then((res: any) => {
              if (res.status === 200) {
                mqttSuccessCnt ++;
                if (mqttSuccessCnt === mqttTotalCnt) {
                  this.northInfoLoading = false;
                  if (tipType === 'mqtt') {
                    this.$Message.info('保存成功');
                    this.northInfoDisabled = true;
                  }
                  this.saveInfoToLocal();
                }
              }
            });
          } else {
            addMQTTNorthboundProtocolApi(this.getNorthMqttParam(mqttItem)).then((res: any) => {
              if (res.status === 200) {
                mqttItem.uuid = res.data.uuid;
                mqttSuccessCnt ++;
                if (mqttSuccessCnt === mqttTotalCnt) {
                  this.northInfoLoading = false;
                  if (tipType === 'mqtt') {
                    this.$Message.info('保存成功');
                    this.northInfoDisabled = true;
                  }
                  this.saveInfoToLocal();
                }
              }
            });
          }
        } else {
          this.$Message.error('MQTT 参数错误');
          this.northInfoLoading = false;
        }
      });

      // tibco
      this.northInfo.tibcoRvDataList.forEach((tibcoItem: any, index: any) => {
        if (this.checkNorthinfoTibcoAll(tibcoItem, index)) {
          this.northInfoLoading = true;
          if (tibcoItem.uuid) {
            editTibcoRvNorthboundProtocolApi(this.getNorthTibcoRVParam(tibcoItem)).then((res: any) => {
              if (res.status === 200) {
                tibcoSuccessCnt ++;
                if (tibcoSuccessCnt === tibcoTotalCnt) {
                  this.northInfoLoading = false;
                  if (tipType === 'tibcorv') {
                    this.$Message.info('保存成功');
                    this.northInfoDisabled = true;
                  }
                  this.saveInfoToLocal();
                }
              }
            });
          } else {
            addTibcoRvNorthboundProtocolApi(this.getNorthTibcoRVParam(tibcoItem)).then((res: any) => {
              if (res.status === 200) {
                tibcoItem.uuid = res.data.uuid;
                tibcoSuccessCnt ++;
                if (tibcoSuccessCnt === tibcoTotalCnt) {
                  this.northInfoLoading = false;
                  if (tipType === 'tibcorv') {
                    this.$Message.info('保存成功');
                    this.northInfoDisabled = true;
                  }
                  this.saveInfoToLocal();
                }
              }
            });
          }
        } else {
          this.$Message.error('TibcoRV 参数错误');
          this.northInfoLoading = false;
        }
      });
    }

  }

  // 查询 MQTT 北向协议
  private queryMQTTNorthboundProtocol(item: any, isTemplate: boolean) {
    const mqttLength = item.mqttnorthboundList && item.mqttnorthboundList.length;
    if (mqttLength > 0) {
      item.mqttnorthboundList.forEach((id: any, index: any) => {
        queryMQTTNorthboundProtocolApi(id).then((res: any) => {
          const tmpMqtt = isTemplate ? ({...res.data, ...this.northInfo.checkNorthInfoMqttObj, uuid: '', testMqtt: false, mqttServerTopicError: [] }) : ({...res.data, ...this.northInfo.checkNorthInfoMqttObj, testMqtt: false, mqttServerTopicError: []});
          this.northInfo.mqttDataList.push(tmpMqtt);
          const checkIndex = this.northInfo.mqttDataList.length - 1;
          this.saveInfoToLocal();
          this.checkNorthInfoMqttAll(this.northInfo.mqttDataList[checkIndex], checkIndex);
          this.testMqtt(this.northInfo.mqttDataList[checkIndex]);
        });
      });
    }
  }

  // 北向新增 MQTT
  private addNorthMqtt() {
    this.northInfo.mqttDataList.push(JSON.parse(JSON.stringify(this.northInfoMqttDataOrigin)));
    (this as any).$nextTick(() => {
      if (this.northInfo.tibcoRvDataList.length > 1) {
        (this as any).$refs.addWinCon.scrollTop =  (this as any).$refs.addWinCon.scrollTop - (this.northInfo.tibcoRvDataList.length * 150);
      } else {
        (this as any).$refs.addWinCon.scrollTop += 224;
      }
    });
  }

  // 北向删除 MQTT
  private deleteMQTTNorthboundProtocol(item: any, index: any) {
    if (item.uuid) {
      deleteMQTTNorthboundProtocolApi(item.uuid).then((res: any) => {
        if (res.status === 200) {
          this.northInfo.mqttDataList.splice(index, 1);
          this.$Message.info('删除成功');
          this.saveInfoToLocal();
        }
      });
    } else {
      this.northInfo.mqttDataList.splice(index, 1);
      this.saveInfoToLocal();
    }
  }

  // 北向测试 MQTT
  private testMqtt(item: any) {
    testMqttConnectApi(item).then((res: any) => {
      if (res.status === 200) {
        item.testMqtt = true;
        this.$Message.info(res.msg);
      } else {
        item.testMqtt = false;
      }
      this.saveInfoToLocal();
    });
  }

  // 北向查询 TibcoRV
  private queryTibcoRvNorthboundProtocol(item: any, isTemplate: boolean) {
    const tibocoLength = item.tibocoRvnorthboundList && item.tibocoRvnorthboundList.length;
    if ( tibocoLength > 0) {
      item.tibocoRvnorthboundList.forEach((id: any, index: any) => {
        queryTibcoRvNorthboundProtocolApi(id).then((res: any) => {
          const tmpTibco = isTemplate ? ({...res.data, ...this.northInfo.checkNorthInfoTibcoObj, uuid: ''}) : ({...res.data, ...this.northInfo.checkNorthInfoTibcoObj});
          this.northInfo.tibcoRvDataList.push(tmpTibco);
          const checkIndex = this.northInfo.tibcoRvDataList.length - 1;
          this.checkNorthinfoTibcoAll(this.northInfo.tibcoRvDataList[checkIndex], checkIndex);
          this.saveInfoToLocal();
        });
      });
    }
  }

  // 北向新增 TibcoRV
  private addNorthTibco() {
    this.northInfo.tibcoRvDataList.push(JSON.parse(JSON.stringify(this.northInfoTibcoDataOrigin)));
    (this as any).$nextTick(() => {
      (this as any).$refs.addWinCon.scrollTop += 150;
    });
  }

  // 北向删除 TibcoRV
  private deleteTibcoRvNorthboundProtocol(item: any, index: any) {
    if (item.uuid) {
      deleteTibcoRvNorthboundProtocolApi(item.uuid).then((res: any) => {
        if (res.status === 200) {
          this.northInfo.tibcoRvDataList.splice(index, 1);
          this.$Message.info('删除成功');
          this.saveInfoToLocal();
        }
      });
    } else {
      this.northInfo.tibcoRvDataList.splice(index, 1);
      this.saveInfoToLocal();
    }
  }

  // 北向取消删除 TibcoRV
  private cancelDeleteTibcoRvNorthboundProtocol() {
    //
  }

  // 北向校验单个 MQTT 参数
  private checkNorthInfoMqttItem(item: any, type: any, index: any) {
    let isValid: boolean = true;
    const reg = /^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/;
    if (type === 'mqttServerIp') {
      if (item.mqttServerIp && reg.test(item.mqttServerIp)) {
        item[`${type}Error`] = false;
      } else {
        item[`${type}Error`] = true;
        isValid = false;
      }
    } else if (type === 'mqttServerPort') {
      if (item.mqttServerPort > 0 && item.mqttServerPort <= 65535) {
        item[`${type}Error`] = false;
      } else {
        item[`${type}Error`] = true;
        isValid = false;
      }
    } else if (type === 'topic') {
      item.mqttServerTopic.forEach((topic: any, mqttIndex: any) => {
        if (topic) {
          this.$set(item.mqttServerTopicError, mqttIndex, false);
        } else {
          this.$set(item.mqttServerTopicError, mqttIndex, true);
          isValid = false;
        }
      });
    } else {
      if (item[type]) {
        item[`${type}Error`] = false;
      } else {
        item[`${type}Error`] = true;
        isValid = false;
      }
    }

    return isValid;
  }
  // 北向校验全部 MQTT 参数
  private checkNorthInfoMqttAll(mqttItem: any, index: any) {
    let isValid: boolean = true;
    const checkNorthInfoMqttList = ['mqttServerIp', 'mqttServerPort', 'mqttServerClientID', 'topic'];
    checkNorthInfoMqttList.forEach((type: any) => {
      if (!this.checkNorthInfoMqttItem(mqttItem, type, index)) {
        isValid = false;
      }
    });

    return isValid;
  }
  // 北向校验单个 TibcoRV
  private checkNorthInfoTibcoItem(item: any, type: any, index: any) {
    let isValid: boolean = true;
    if (item[type]) {
      item[`${type}Error`] = false;
    } else {
      item[`${type}Error`] = true;
      isValid = false;
    }

    return isValid;
  }
  // 北向校验全部 TibcoRV
  private checkNorthinfoTibcoAll(tibcoItem: any, index: any) {
    let isValid: boolean = true;
    const checkNorthInfoTibcoList = ['service', 'network', 'daemon', 'subject'];
    checkNorthInfoTibcoList.forEach((type: any) => {
      if (!this.checkNorthInfoTibcoItem(tibcoItem, type, index)) {
        isValid = false;
      }
    });

    return isValid;
  }


  // 北向显示或隐藏新增类型 Poptip
  private toggleAddNorthPoptip() {
    this.northInfo.isShowAddNorthPoptip ? this.hide() : this.show();
  }
  // hidePanel
  private hidePanel(e: any) {
    if (!(this as any).$refs.poptipContent.contains(e.target)) {
      this.hide();
    }
  }
  // hide
  private hide() {
    this.northInfo.isShowAddNorthPoptip = false;
    document.removeEventListener('click', this.hidePanel, false);
  }
  // show
  private show() {
    this.northInfo.isShowAddNorthPoptip = true;
    document.addEventListener('click', this.hidePanel, false);
  }
}
</script>

<style lang="less" scoped>
@font-color: #333333;
@border-color: #888888;
.title {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  border-bottom: 1px solid @border-color;
  padding-bottom: 9px;
  .name {
    display: inline-block;
    font-size: 16px;
  }
  .info {
    .remark {
      color: @border-color;
    }
  }
}
.circle {
  display: inline-block;
  width:14px;
  height:14px;
  background: #888;
  border:1px solid rgba(136, 136, 136, 1);
  border-radius:7px;
  margin-right: 3px;
  vertical-align: sub;
  &.green {
    background:rgba(18,239,13,1);
  }
  &.yellow {
    background: #efd70d;
  }
}
.child-device {
  .quick-set-config {
    padding-top: 30px;
    .con {
      padding: 30px 100px 50px 100px;
      .con-tip {
        padding-bottom: 32px;
        border-bottom: 1px solid #888;
        margin-bottom: 30px;
      }
      .export-import {
        display: inline-block;
        display: flex;
        justify-content: flex-start;
        .con-export {
          padding-right: 100px;
          .a {
            color: @font-color;
          }
        }
        .con-import {
          .con-import-detail {
            display: flex;
            .name {
              display: inline-block;
              width: 180px;
              height: 32px;
              line-height: 30px;
              background: #fff;
              padding: 0 15px;
              margin-right: 10px;
              border: 1px solid #888888;
              border-radius: 3px;
            }
            .file-upload {
              display: inline-block;
              position: relative;
              width: 90px;
              height: 32px;
              .file-input {
                width: 100%;
                height: 100%;
                position: absolute;
                left: 0;
                opacity: 0;
              }
              .file-btn {
                position: absolute;
                left: 0;
              }
            }
            .error-tip {
              line-height: 32px;
              margin-left: 30px;
              color: #F4333C;
              .iconfont {
                padding-right: 8px;
              }
            }
          }
        }
      }
    }
  }
  .child-set-config {
    .con {
      padding: 30px 100px 36px;
      .device-list {
        display: flex;
        justify-content: flex-start;
        flex-wrap: wrap;
        border-bottom: 1px solid #888;
        .device-item {
          width: 140px;
          height: 58px;
          box-shadow: 0px 1px 3px 0px rgba(136,136,136,0.59);
          border-radius: 3px;
          background: #fff;
          padding: 10px;
          margin-bottom: 36px;
          cursor: pointer;
          &:not(:last-child) {
            margin-right: 20px;
          }
          &.selected {
            border:1px solid rgba(77, 157, 237, 1);
            box-shadow:0px 1px 3px 0px rgba(77,157,237,0.59);
          }
          .name {
            text-align: right;
            .icon-edit {
              padding-right: 10px;
              color: #FF5B05;
            }
          }
        }
      }

      .com-wrap {
        border-bottom: 1px solid #888;
        .com-con {
          text-align: left;
          .com-name {
            padding: 16px 0 24px 0;
          }
          .com-set {
            .com-set-item {
              display: inline-block;
              padding-bottom: 38px;
              &:not(:last-child) {
                margin-right: 30px;
              }
            }
          }
        }
      }

      .option-wrap {
        width: 100%;
        &.fix {
          z-index: 10;
          position: fixed;
          top: 40px;
          right: 0;
          left: 200px;
          padding: 0px 130px;
          width: calc(100% - 200px);
          box-shadow: rgba(0, 0, 0, 0.1) 0px 1px 0px;
          background-color: #ECECEE;
        }
        .option-con {
          display: flex;
          justify-content: space-between;
          padding: 30px 0;
        }
      }

      .device-detail-wrap {
        min-width: 815px;
        .device-detail-item {
          position: relative;
          background: #fff;
          height:90px;
          box-shadow:0px 0px 3px 0px rgba(136,136,136,0.59);
          border-radius:3px;
          margin-bottom: 25px;
          padding: 15px;
          font-weight: 400;
          color: #333;
          cursor: pointer;
          &.selected {
            border:1px solid rgba(77, 157, 237, 1);
          }
          .message {
            line-height: 36px;
            .name {
              font-size: 16px;
              margin-right: 5px;
              display: inline-block;
              width: 130px;
              .icon-edit {
                color: #FF5B05;
                padding-left: 1px;
                margin-right: -25px;
              }
            }
            .message-item {
              margin-right: 30px;
              color: #888;
              .circle {
                .circle;
              }
            }
          }
          .nature {
            margin-left: 130px;
            .nature-item {
              margin-right: 30px;
            }
          }
          .option {
            position: absolute;
            right: 0;
            top: 50%;
            transform: translateY(-50%);
            line-height: 60px;
            padding: 0 24px 0 21px;
            border-left: 1px solid #888;
          }
        }
      }
    }
  }

  .delete-win {
    .delete-win-con {
      padding: 20px 30px 20px 50px;
      .delete-device-list {
        display: flex;
        justify-content: flex-start;
        flex-wrap: wrap;
        .delete-device-item {
          padding: 3px 20px;
          border: 1px solid rgba(230, 230, 232, 1);
          border-radius: 3px;
          margin-right: 20px;
          margin-bottom: 20px;
          cursor: pointer;
          &.selected {
            border:1px solid rgba(77, 157, 237, 1);
          }
        }
      }
      .footer {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-top: 10px;
        margin-right: 23px;
      }
    }
  }

  .add-win {
    .add-win-con {
      max-height: 600px;
      overflow-y: auto;
      padding: 24px 30px;
      .error-text {
        position: absolute;
        left: 1px;
        top: 27px;
        color: #F4333C;
      }
      .footer {
        text-align: right;
        margin: 20px 20px 56px 0;
      }
      .child-device-info {
        .title {
          .title;
        }
        .content {
          text-align: left;
          .common {
            margin: 22px 0 0 22px;
          }
          .variable {
            background:#ECECEE;
            padding: 20px;
            border-radius: 5px;
            .inline-item {
              .ivu-form-item {
                display: inline-block;
              }
            }
            .block-item {
              margin-bottom: -20px;
            }
          }
        }
      }
      .south {
        .title {
          .title;
        }
        .content {
          text-align: left;
          .treaty-wrap {
            margin: 20px;
            border-bottom: 1px solid @border-color;
          }
          .time-wrap {
            margin: 20px;
          }
          .variable {
            background:#ECECEE;
            padding: 20px;
            border-radius: 5px;
            .inline-item {
              .ivu-form-item {
                display: inline-block;
                margin-right: 4px;
                .ivu-form-item-label {
                  padding-right: 0;
                }
              }
            }
            .variable-item {
              border-top: 1px solid @border-color;
              padding-top: 20px;
              .row {
                &:after {
                  content: '';
                  display: block;
                  clear: both;
                }
                .ivu-form-item {
                  display: inline-block;
                  margin-right: 8px;
                }
                .btn {
                  cursor: pointer;
                  float: right;
                  .iconfont {
                    font-size: 20px;
                    &:last-child {
                      margin-left: 20px;
                    }
                  }
                }
              }
            }
          }
        }
      }
      .north {
        margin-bottom: 50px;
        .title {
          .title;
        }
        .content {
          .mqtt-wrap, .tibco-wrap {
            background: #ECECEE;
            padding: 20px;
            border-radius: 5px;
            margin-top: 20px;
            .mqtt-item, .tibco-item {
              &:not(:first-child) {
                border-top: 1px solid @border-color;
                padding-top: 20px;
              }
              .option {
                display: flex;
                justify-content: space-between;
                align-items: center;
                margin-bottom: 20px;
              }
              .btn {
                cursor: pointer;
                .test-mqtt {
                  display: inline-block;
                  padding-right: 30px;
                }
              }
              .info {
                text-align: left;
                .row {
                  .ivu-form-item {
                    display: inline-block;
                    margin-right: 22px;
                    &:last-child {
                      margin-right: 0;
                    }
                  }
                  &.row-bottom {
                    border-bottom: 1px solid #888;
                    margin-bottom: 20px;
                  }
                  .topic-btn {
                    display: inline-block;
                    line-height: 32px;
                    cursor: pointer;
                    .icon-plus {
                      padding-right: 18px;
                    }
                  }
                }
                .topic-list {
                  .topic-has-axle-wrap {
                    position: absolute;
                    width: 70px;
                    left: -75px;
                    &.hirata {
                      width: 60px;
                      left: -60px;
                    }
                    .topic-has-axle-container {
                      text-align: center;
                      display: inline-block;
                      display: flex;
                      flex-direction: column;
                      justify-content: flex-start;
                      // align-items: flex-start;
                      align-items: center;
                      .topic-has-axle {
                        line-height: 16px;
                      }
                      .topic-axle {
                        line-height: 14px;
                        color: #999;
                      }
                    }
                  }
                }
              }
              
            }
          }
        }
        .north-footer {
          display: flex;
          justify-content: space-between;
          margin: 30px 20px 56px 0;
          .poptip {
            position: relative;
            .tip {
              width: 84px;
              position: absolute;
              top: -9px;
              right: -100px;
              box-shadow: 0px 0px 3px 0px rgba(136,136,136,0.59);
              border-radius: 3px;
              cursor: pointer;
              &::before {
                content: '';
                position: absolute;
                top: 17px;
                left: -10px;
                width: 0;
                height: 0;
                border-right: 10px solid #E4E4E2;
                border-bottom: 10px solid transparent;
                border-top: 10px solid transparent; 
              }
              &::after {
                content: '';
                position: absolute;
                top: 18px;
                left: -9px;
                width: 0;
                height: 0;
                border-bottom: 9px solid transparent;
                border-right: 9px solid #fff;
                border-top: 9px solid transparent; 
              }
              p {
                padding: 4px 0;
                &:hover {
                  background: #A6CEF6;
                }
              }
            }
          }
        }
      }
    }
  }

  .close-confirm-win {
    .close-confirm-con {
      padding: 30px 50px 20px 50px;
      .close-tip {
        padding-bottom: 42px;
      }
      .btn-cancel {
        margin-left: 14px;
      }
    }
  }
}
</style>
<style lang="less">
.north-footer {
  .ivu-poptip-popper {
    min-width: 101px;
    .ivu-poptip-body {
      padding: 0;
    }
  }
}
</style>
