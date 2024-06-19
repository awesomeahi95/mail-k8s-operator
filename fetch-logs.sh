for ns in $(kubectl get namespaces -o jsonpath="{.items[*].metadata.name}"); do
  echo "Namespace: $ns"
  for pod in $(kubectl get pods -n $ns -l control-plane=controller-manager -o jsonpath="{.items[*].metadata.name}"); do
    echo "Pod: $pod"
    kubectl logs -n $ns $pod
  done
done
