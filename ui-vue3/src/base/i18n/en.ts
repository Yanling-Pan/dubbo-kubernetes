/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import type { I18nType } from './type.ts'

const words: I18nType = {
  service: 'Service',
  serviceSearch: 'Search Service',
  serviceGovernance: 'Routing Rule',
  trafficManagement: 'Traffic Management',
  routingRule: 'Condition Rule',
  tagRule: 'Tag Rule',
  meshRule: 'Mesh Rule',
  dynamicConfig: 'Dynamic Config',
  accessControl: 'Black White List',
  weightAdjust: 'Weight Adjust',
  loadBalance: 'Load Balance',
  serviceTest: 'Service Test',
  serviceMock: 'Service Mock',
  serviceMetrics: 'Service Metrics',
  serviceRelation: 'Service Relation',
  metrics: 'Metrics',
  relation: 'Relation',
  group: 'Group',
  serviceInfo: 'Service Info',
  providers: 'Providers',
  consumers: 'Consumers',
  common: 'Common',
  version: 'Version',
  app: 'Application',
  services: 'Services',
  application: 'Application',
  instance: 'Instance',
  all: 'All',
  ip: 'IP',
  qps: 'qps',
  rt: 'rt',
  successRate: 'success rate',
  port: 'PORT',
  timeout: 'timeout(ms)',
  serialization: 'serialization',
  appName: 'Application Name',
  serviceName: 'Service Name',
  registrySource: 'Registry Source',
  instanceRegistry: 'Instance Registry',
  interfaceRegistry: 'Interface Registry',
  allRegistry: 'Instance / Interface Registry',
  operation: 'Operation',
  searchResult: 'Search Result',
  search: 'Search',
  methodName: 'Method Name',
  enabled: 'Enabled',
  disabled: 'Disabled',
  method: 'Method',
  weight: 'Weight',
  create: 'CREATE',
  save: 'SAVE',
  cancel: 'CANCEL',
  close: 'CLOSE',
  confirm: 'CONFIRM',
  ruleContent: 'RULE CONTENT',
  createNewRoutingRule: 'Create New Routing Rule',
  createNewTagRule: 'Create New Tag Rule',
  createNewMeshRule: 'Create New Mesh Rule',
  createNewDynamicConfigRule: 'Create New Dynamic Config Rule',
  createNewWeightRule: 'Create New Weight Rule',
  createNewLoadBalanceRule: 'Create new load balancing rule',
  createTimeoutRule: 'Create timeout rule',
  createRetryRule: 'Create timeout rule',
  createRegionRule: 'Create retry rule',
  createArgumentRule: 'Create argument routing rule',
  createMockCircuitRule: 'Create mock (circuit breaking) rule',
  createAccesslogRule: 'Create accesslog rule',
  createGrayRule: 'Create gray rule',
  createWeightRule: 'Create weighting rule',
  serviceIdHint: 'Service ID',
  view: 'View',
  edit: 'Edit',
  delete: 'Delete',
  searchRoutingRule: 'Search Routing Rule',
  searchAccess: 'Search Access Rule',
  searchWeightRule: 'Search Weight Adjust Rule',
  dataIdClassHint: 'Complete package path of service interface class',
  dataIdVersionHint:
    'The version of the service interface, which can be filled in according to the actual situation of the interface',
  dataIdGroupHint:
    'The group of the service interface, which can be filled in according to the actual situation of the interface',
  agree: 'Agree',
  disagree: 'Disagree',
  searchDynamicConfig: 'Search Dynamic Config',
  appNameHint: 'Application name the service belongs to',
  basicInfo: 'BasicInfo',
  metaData: 'MetaData',
  methodMetrics: 'Method Statistics',
  searchDubboService: 'Search Dubbo Services or applications',
  serviceSearchHint: 'Service ID, org.apache.dubbo.demo.api.DemoService, * for all services',
  ipSearchHint: 'Find all services provided by the target server on the specified IP address',
  appSearchHint:
    'Input an application name to find all services provided by one particular application, * for all',
  searchTagRule: 'Search Tag Rule by application name',
  searchMeshRule: 'Search Mesh Rule by application name',
  searchSingleMetrics: 'Search Metrics by IP',
  searchBalanceRule: 'Search Balancing Rule',
  noMetadataHint:
    'There is no metadata available, please update to Dubbo2.7, or check your config center configuration in application.properties, please check ',
  parameterList: 'parameterList',
  returnType: 'returnType',
  here: 'here',
  configAddress: 'https://github.com/apache/incubator-dubbo-admin/wiki/Dubbo-Admin-configuration',
  whiteList: 'White List',
  whiteListHint: 'White list IP address, divided by comma: 1.1.1.1,2.2.2.2',
  blackList: 'Black List',
  blackListHint: 'Black list IP address, divided by comma: 3.3.3.3,4.4.4.4',
  address: 'Address',
  weightAddressHint: 'IP addresses to set this weight, divided by comma: 1.1.1.1,2.2.2.2',
  weightHint: 'weight value, default is 100',
  methodHint: 'choose method of load balancing, * for all methods',
  strategy: 'Strategy',
  balanceStrategyHint: 'load balancing strategy',
  goIndex: 'Go To Index',
  releaseLater: 'will release later',
  later: {
    metrics: 'Metrics will release later',
    serviceTest: 'Service Test will release later',
    serviceMock: 'Service Mock will release later'
  },
  by: 'by ',
  $vuetify: {
    dataIterator: {
      rowsPerPageText: 'Items per page:',
      rowsPerPageAll: 'All',
      pageText: '{0}-{1} of {2}',
      noResultsText: 'No matching records found',
      nextPage: 'Next page',
      prevPage: 'Previous page'
    },
    dataTable: {
      rowsPerPageText: 'Rows per page:'
    },
    noDataText: 'No data available'
  },
  configManage: 'Configuration Management',
  configCenterAddress: 'ConfigCenter Address',
  searchDubboConfig: 'Search Dubbo Config',
  createNewDubboConfig: 'Create New Dubbo Config',
  scope: 'Scope',
  name: 'Name',
  warnDeleteConfig: ' Are you sure to Delete Dubbo Config: ',
  warnDeleteRouteRule: 'Are you sure to Delete routing rule',
  warnDeleteDynamicConfig: 'Are you sure to Delete dynamic config',
  warnDeleteBalancing: 'Are you sure to Delete load balancing',
  warnDeleteAccessControl: 'Are you sure to Delete access control',
  warnDeleteTagRule: 'Are you sure to Delete tag rule',
  warnDeleteMeshRule: 'Are you sure to Delete mesh rule',
  warnDeleteWeightAdjust: 'Are you sure to Delete weight adjust',
  configNameHint:
    "Application name the config belongs to, use 'global'(without quotes) for global config",
  configContent: 'Config Content',
  testMethod: 'Test Method',
  execute: 'EXECUTE',
  result: 'Result: ',
  success: 'SUCCESS',
  fail: 'FAIL',
  detail: 'Detail',
  more: 'More',
  copyUrl: 'Copy URL',
  copy: 'Copy',
  url: 'URL',
  copySuccessfully: 'Copied',
  test: 'Test',
  placeholders: {
    searchService: 'Search by service name'
  },
  methods: 'Methods',
  testModule: {
    searchServiceHint:
      'Entire service ID, org.apache.dubbo.demo.api.DemoService, press Enter to search'
  },
  userName: 'User Name',
  password: 'Password',
  login: 'Login',
  apiDocs: 'API Docs',
  apiDocsRes: {
    dubboProviderIP: 'Dubbo Provider Ip',
    dubboProviderPort: 'Dubbo Provider Port',
    loadApiList: 'Load Api List',
    apiListText: 'Api List',
    apiForm: {
      missingInterfaceInfo: 'Missing interface information',
      getApiInfoErr: 'Exception in obtaining interface information',
      api404Err:
        'Interface name is incorrect, interface parameters and response information are not found',
      apiRespDecShowLabel: 'Response Description',
      apiNameShowLabel: 'Api Name',
      apiPathShowLabel: 'Api Path',
      apiMethodParamInfoLabel: 'Api method parameters',
      apiVersionShowLabel: 'Api Version',
      apiGroupShowLabel: 'Api Group',
      apiDescriptionShowLabel: 'Api Description',
      isAsyncFormLabel:
        'Whether to call asynchronously (this parameter cannot be modified, according to whether to display asynchronously defined by the interface)',
      apiModuleFormLabel: 'Api module (this parameter cannot be modified)',
      apiFunctionNameFormLabel: 'Api function name(this parameter cannot be modified)',
      registryCenterUrlFormLabel:
        'Registry address. If it is empty, Dubbo provider IP and port will be used for direct connection',
      paramNameLabel: 'Parameter name',
      paramPathLabel: 'Parameter path',
      paramDescriptionLabel: 'Description',
      paramRequiredLabel: 'This parameter is required',
      doTestBtn: 'Do Test',
      responseLabel: 'Response',
      responseExampleLabel: 'Response Example',
      apiResponseLabel: 'Api Response',
      LoadingLabel: 'Loading...',
      requireTip: 'There are required items not filled in',
      requireItemTip: 'This field is required',
      requestApiErrorTip:
        'There is an exception in the request interface. Please check the submitted data, especially the JSON class data and the enumeration part',
      unsupportedHtmlTypeTip: 'Temporarily unsupported form type',
      none: 'none'
    }
  },
  authFailed: 'Authorized failed,please login.',

  ruleList: 'Rule List',
  mockRule: 'Mock Rule',
  mockData: 'Mock Data',
  globalDisable: 'Global Disable',
  globalEnable: 'Global Enable',
  saveRuleSuccess: 'Save Rule Successfully',
  deleteRuleSuccess: 'Delete Rule Successfully',
  disableRuleSuccess: 'Disable Rule Successfully',
  enableRuleSuccess: 'Enable Rule Successfully',
  methodNameHint: 'The method name of Service',
  createMockRule: 'Create Mock Rule',
  editMockRule: 'Edit Mock Rule',
  deleteRuleTitle: 'Are you sure to delete this mock rule?',

  trafficTimeout: 'Timeout',
  trafficRetry: 'Retry',
  trafficRegion: 'Region Aware',
  trafficIsolation: 'Isolation',
  trafficWeight: 'Weight Percentage',
  trafficArguments: 'Arg Routing',
  trafficMock: 'Mock',
  trafficAccesslog: 'Accesslog',
  trafficHost: 'Host',
  homePage: 'Cluster Overview',
  serviceManagement: 'Dev & Test',
  resources: 'Resources',
  applications: 'Applications',
  instances: 'Instances',
  applicationDomain: {
    name: 'Application Name',
    detail: 'Application Detail'
  },
  searchDomain: {
    total: 'Total',
    unit: 'items'
  },
  backHome: 'Back Home',
  noPageTip: 'Sorry, the page you visited does not exist.',
  globalSearchTip: 'Search ip, application, instance, service'
}

export default words
