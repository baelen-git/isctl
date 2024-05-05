
#this is the name used for created test objects. Must match the names in YAML files in tests/data/...
TEST_NAME=isctl-bats-test   

TEST_SECTION="Apply"

@test "${TEST_SECTION}: apply all-in-one YAML file" {
    ./build/isctl ${ISCTL_OPTIONS} apply -f tests/data/apply-yaml-all-in-one/test.yaml
}

@test "${TEST_SECTION}: get NTP policy list should include new policy" {
    ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${TEST_NAME}"
    ./build/isctl ${ISCTL_OPTIONS} get organization organization | grep "${TEST_NAME}"
}

@test "${TEST_SECTION}: get NTP policy by name and check NTP servers correct" {
    NTPSERVERS=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NAME}" -o json | jq -r -c .NtpServers )
    
    [ "${NTPSERVERS}" == "[\"1.1.1.1\"]" ]
}

@test "${TEST_SECTION}: apply directoy of YAML files" {
    ./build/isctl ${ISCTL_OPTIONS} apply -f tests/data/apply-yaml
}

@test "${TEST_SECTION}: check updated NTP servers correct" {
    NTPSERVERS=$( ./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NAME}" -o json | jq -r -c .NtpServers )
    
    [ "${NTPSERVERS}" == "[\"1.1.1.1\",\"2.2.2.2\"]" ]
}

@test "${TEST_SECTION}: delete NTP policy" {
    ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NAME}" -o jsonpath='$.Moid')
    sleep 5
    ./build/isctl ${ISCTL_OPTIONS} delete organization organization moid $(./build/isctl ${ISCTL_OPTIONS} get organization organization --name "${TEST_NAME}" -o jsonpath='$.Moid')

    # check that the policy no longer exists
    ! ./build/isctl ${ISCTL_OPTIONS} get ntp policy | grep "${TEST_NAME}"
    ! ./build/isctl ${ISCTL_OPTIONS} get organization organization | grep "${TEST_NAME}"
}

@test "${TEST_SECTION}: bulk.MoCloners" {
    ./build/isctl ${ISCTL_OPTIONS} create server profiletemplate --Name "${TEST_NAME}" --Organization default

    ./build/isctl ${ISCTL_OPTIONS} apply -f tests/data/test-bulk-mo-cloner.yaml

    run ./build/isctl ${ISCTL_OPTIONS} delete server profile name "${TEST_NAME}-1"
    run ./build/isctl ${ISCTL_OPTIONS} delete server profiletemplate name "${TEST_NAME}" 
}

@test "${TEST_SECTION}: apply for MO with no Name attribute" {
    run ./build/isctl ${ISCTL_OPTIONS} apply -f tests/data/test-appliance-device-claim.yaml
    # this should fail because we're are testing against SaaS, not an appliance but as long as the error code matches we know the request was sent properly
    assert_failure
    assert_line --partial "Performing create operation on new MO (ClassId: appliance.DeviceClaim)"
    assert_line --partial "Error while applying MOs: error executing operation: request failed: 403 Forbidden: Operation not supported."
}

@test "${TEST_SECTION}: create and delete YAML file with unnamed objects" {
    ./build/isctl ${ISCTL_OPTIONS} apply -f tests/data/test-delete-with-unnamed.yaml

    run ./build/isctl get fabric portpolicy name isctl-bats-test
    assert_success

    ./build/isctl ${ISCTL_OPTIONS} apply -d -f tests/data/test-delete-with-unnamed.yaml

    run ./build/isctl get fabric portpolicy name isctl-bats-test
    assert_failure
}

setup() {
    load 'test_helper/bats-support/load' # this is required by bats-assert!
    load 'test_helper/bats-assert/load'
}

setup_file() {
    # delete the test objects if they already exist. Don't check the exit code. 
    run ./build/isctl ${ISCTL_OPTIONS} delete ntp policy moid $(./build/isctl ${ISCTL_OPTIONS} get ntp policy --name "${TEST_NAME}" -o jsonpath='$.Moid')
    run ./build/isctl ${ISCTL_OPTIONS} delete organization organization moid $(./build/isctl ${ISCTL_OPTIONS} get organization organization --name "${TEST_NAME}" -o jsonpath='$.Moid')
}