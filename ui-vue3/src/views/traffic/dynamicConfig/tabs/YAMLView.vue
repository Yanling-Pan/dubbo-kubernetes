<!--
  ~ Licensed to the Apache Software Foundation (ASF) under one or more
  ~ contributor license agreements.  See the NOTICE file distributed with
  ~ this work for additional information regarding copyright ownership.
  ~ The ASF licenses this file to You under the Apache License, Version 2.0
  ~ (the "License"); you may not use this file except in compliance with
  ~ the License.  You may obtain a copy of the License at
  ~
  ~     http://www.apache.org/licenses/LICENSE-2.0
  ~
  ~ Unless required by applicable law or agreed to in writing, software
  ~ distributed under the License is distributed on an "AS IS" BASIS,
  ~ WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  ~ See the License for the specific language governing permissions and
  ~ limitations under the License.
-->

<template>
  <a-card>
    <a-flex style="width: 100%">
      <a-col :span="isDrawerOpened ? 24 - sliderSpan : 24" class="left">
        <a-flex vertical align="end">
          <a-button type="text" style="color: #0a90d5" @click="isDrawerOpened = !isDrawerOpened">
            {{ $t('flowControlDomain.versionRecords') }}
            <DoubleLeftOutlined v-if="!isDrawerOpened" />
            <DoubleRightOutlined v-else />
          </a-button>

          <div class="editorBox">
            <MonacoEditor
              :modelValue="YAMLValue"
              theme="vs-dark"
              :height="500"
              language="yaml"
              :readonly="!isEdit"
            />
          </div>
        </a-flex>
      </a-col>

      <a-col :span="isDrawerOpened ? sliderSpan : 0" class="right">
        <a-card v-if="isDrawerOpened" class="sliderBox">
          <a-card v-for="i in 2" :key="i">
            <p>修改时间: 2024/3/20 15:20:31</p>
            <p>版本号: xo842xqpx834</p>

            <a-flex justify="flex-end">
              <a-button type="text" style="color: #0a90d5">查看</a-button>
              <a-button type="text" style="color: #0a90d5">回滚</a-button>
            </a-flex>
          </a-card>
        </a-card>
      </a-col>
    </a-flex>
  </a-card>
  <a-flex v-if="isEdit" style="margin-top: 30px">
    <a-button type="primary">确认</a-button>
    <a-button style="margin-left: 30px">取消</a-button>
  </a-flex>
</template>

<script setup lang="ts">
import MonacoEditor from '@/components/editor/MonacoEditor.vue'
import { DoubleLeftOutlined, DoubleRightOutlined } from '@ant-design/icons-vue'
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const isEdit = ref(route.params.isEdit === '1')
const isDrawerOpened = ref(false)

const sliderSpan = ref(8)

const YAMLValue = ref(
  'configVersion: v3.0\n' +
    'force: true\n' +
    'enabled: true\n' +
    'key: shop-detail\n' +
    'tags:\n' +
    '  - name: gray\n' +
    '    match:\n' +
    '      - key: env\n' +
    '        value:\n' +
    '          exact: gray'
)
</script>

<style scoped lang="less">
.editorBox {
  border-radius: 0.3rem;
  overflow: hidden;
  width: 100%;
}

.sliderBox {
  margin-left: 5px;
  max-height: 530px;
  overflow: auto;
}

&:deep(.left.ant-col) {
  transition: all 0.5s ease;
}

&:deep(.right.ant-col) {
  transition: all 0.5s ease;
}
</style>
