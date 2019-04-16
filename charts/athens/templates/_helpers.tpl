{{- define "fullname" -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- define "readinessPath" -}}
{{- if contains "v0.2.0" .Values.image.tag -}}/{{- else -}}/readyz{{- end -}}
{{- end -}}
{{- define "censusai" -}}
- name: {{ template "fullname" . }}-censusai
  image: {{ .Values.censusSidecarImage }}
  env:
    - name: APPINSIGHTS_INSTRUMENTATIONKEY
    # CENSUSAI_INSTRUMENTATION_KEY
    value: {{ .Values.censusInstrumentationKey }}
{{- end -}}

