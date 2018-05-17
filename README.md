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
- It is required to have at least one test called "test.sh"
- When a custom-resource of type `smoketest` is created, the controller runs the smoke test. The output of the test is stored in the custom-resource.
- If some test other than test.sh needs to be run, the smoketest CR should have the `testToRun` annotation with the name of the test (See `examples/testAnnotation.yaml` as an example)

Add additional notes about how to deploy this on a live system

## Built With

* [Operator-sdk](https://github.com/operator-framework/operator-sdk) - Operator SDK for Kubernetes

