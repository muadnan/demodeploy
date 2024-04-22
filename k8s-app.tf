resource "kubectl_manifest" "argocd_application_loki" {
  validate_schema   = false
  server_side_apply = false
  yaml_body         = <<YAML
    apiVersion: argoproj.io/v1alpha1
    kind: Application
    metadata:
      name: loki
      namespace: ${var.argocd_namespace}
      finalizers:
        - resources-finalizer.argocd.argoproj.io
    spec:
      destination:
        namespace: ${local.namespace}
        server: https://kubernetes.default.svc
      project: default
      syncPolicy:
        automated:
          prune: true
          selfHeal: true
        syncOptions:
          - ServerSideApply=true
      source:
        repoURL: https://grafana.github.io/helm-charts
        chart: loki
        targetRevision: 5.41.4
        helm:
          values: |-
            ${indent(12, local.loki_values)}
  YAML
}
