# Smoke-test-operator

Smoke-test-operator is a Kubernetes operator for running smoke tests against a Kubernetes cluster. These tests can show the current health of various components of the Kubernetes cluster.

## Motivation

Every user and every tenant needs to know the state of a Kubernetes cluster. Currently, each of these use their own mechanism of for finding the health of the cluster. This can be error prone.

A better approach would be if a standard set of tests is made available for everyone to run at anytime. Smoke-test-operator does that by providing a new custom resource called `SmokeTest` and having precreated set of smoke tests in a config map. Each time a SmokeTest resource is created, one test from the test suite is run and it's output is stored as part of the SmokeTest object itself in Kubernetes.

Typically, cluster provisioners would create the set of tests and make them available to the users. Whenever a user needs to know the cluster health, s/he simply creates a SmokeTest CR which runs the test and stores the output. Cluster provisioners can keep adding/modifying exsiting tests without having to inform other users about it.

## Getting Started

The following instructions will get you the smoke-test-operator running on the Kubernetes cluster.

### Prerequisites

Have a Kubernetes cluster installed.

### Installing

Install the custom resource and the controller for it.

```
$ kubectl apply -f deploy/install.yaml
```

Ensure that the "smoketests" custom resource is available

```
$ kubectl describe customresourcedefinition smoketests.smoketest.k8s.io
Name:         smoketests.smoketest.k8s.io
Namespace:
Labels:       <none>
...

$ kubectl get smoketests
No resources found.
```

## Running the tests

The repo has a couple of examples of the smoke tests.

```
$ kubectl create -f examples/basicTest.yaml
smoketest "smoke-test-kllwv" created
```

Check the newly created custom resource (`smoke-test-kllwv` in the above example).
```
$ kubectl get smoketest smoke-test-kllwv -o yaml
...
status:
  TestOutput: |
    Get pods:
    ---------------
    NAME                                   READY     STATUS    RESTARTS   AGE       IP              NODE
    smoke-test-operator-65b64b6fc9-hft4z   1/1       Running   0          2m        100.97.50.124   ip-172-20-34-119.us-west-2.compute.internal
...
```

## Details

- The smoke tests to run are configured in a config-map called `smoke-test-config`
- Actual tests are shell scripts
- The CustomResource called smoketest has a 'testToRun' field in the spec. This field should be set to the test that should be run.
- When the custom-resource is created, the controller reads the spec and runs the smoke test mentioned as 'testToRun'. The output of the test is stored in the custom-resource's Status field.
- The output format can be plain text or json. The Spec also has a field called 'outputFormat'. The default is 'text'. Otherwise, set the 'outputFormat' field to 'json'.

## Built With

* [Operator-sdk](https://github.com/operator-framework/operator-sdk) - Operator SDK for Kubernetes

