{{/*
Return Admin Labels to use.
*/}}
{{- define "admin.labels" -}}
app.kubernetes.io/name: {{ template "admin.name" . }}
helm.sh/chart: {{ include "admin.name" . }}-{{ .Values.image.tag }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Return ZooKeeper Labels to use.
*/}}
{{- define "zoo.labels" -}}
app.kubernetes.io/name: {{ template "zoo.name" . }}
helm.sh/chart: {{ include "zoo.name" . }}-{{ .Values.zookeeper.image.tag }}
app.kubernetes.io/instance: {{ template "zoo.name" . }}
app.kubernetes.io/component: {{ template "zoo.name" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Return Nacos Labels to use.
*/}}
{{- define "nacos.labels" -}}
app.kubernetes.io/name: {{ template "nacos.name" . }}
helm.sh/chart: {{ include "nacos.name" . }}-{{ .Values.nacos.image.tag }}
app.kubernetes.io/instance: {{ template "nacos.name" . }}
app.kubernetes.io/component: {{ template "nacos.name" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Return Job Labels to use.
*/}}
{{- define "job.labels" -}}
app.kubernetes.io/name: {{ template "job.name" . }}
helm.sh/chart: {{ include "job.name" . }}-{{ .Values.jobs.image.tag }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Return Traefik Labels.
*/}}
{{- define "traefik.labels" -}}
app.kubernetes.io/name: {{ template "traefik.name" . }}
{{- end -}}

