{{- define "cores-api-library.service" -}}
apiVersion: v1
kind: Service
metadata:
  name: service-ports
  namespace: {{ .Values.NAMESPACE }}
spec:
  selector:
    app: separate-copy
  ports:
  - name: grpc
    protocol: TCP
    port: {{ .Values.PORTS.INTERNAL_PORT }}
    targetPort: {{ .Values.PORTS.CONTAINER_PORT }}
    nodePort: {{ .Values.PORTS.EXTRENAL_PORT }}
  - name: prometheus
    protocol: TCP
    port: {{ .Values.PORTS.PROMETHEUS_INTERNAL_PORT }}
    targetPort: {{ .Values.PORTS.PROMETHEUS_CONTAINER_PORT }}
    nodePort: {{ .Values.PORTS.PROMETHEUS_EXTRENAL_PORT }}
  type: NodePort
{{- end }}