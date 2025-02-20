"""
Kubernetes CIS rules verification.
This module verifies correctness of retrieved findings by manipulating audit and remediation actions
"""
from datetime import datetime, timedelta
import pytest
from commonlib.utils import get_ES_evaluation

from product.tests.data.file_system import eks_file_system_test_cases as eks_fs_tc
from product.tests.parameters import register_params, Parameters


@pytest.mark.eks_file_system_rules
def test_eks_file_system_configuration(
    elastic_client,
    cloudbeat_agent,
    rule_tag,
    node_hostname,
    expected,
):
    """
    This data driven test verifies rules and findings return by cloudbeat agent.
    In order to add new cases @pytest.mark.parameterize section shall be updated.
    Setup and teardown actions are defined in data method.
    This test creates verifies that cloudbeat returns correct finding.
    @param rule_tag: Name of rule to be verified.
    @param node_hostname: EKS node hostname
    @param expected: Result to be found in finding evaluation field.
    @return: None - Test Pass / Fail result is generated.
    """
    # pylint: disable=duplicate-code

    def identifier(eval_resource):
        try:
            return eval_resource.host.name == node_hostname
        except AttributeError:
            return False

    evaluation = get_ES_evaluation(
        elastic_client=elastic_client,
        timeout=cloudbeat_agent.findings_timeout,
        rule_tag=rule_tag,
        exec_timestamp=datetime.utcnow() - timedelta(hours=1),
        resource_identifier=identifier,
    )

    assert evaluation is not None, f"No evaluation for rule {rule_tag} could be found"
    assert evaluation == expected, f"Rule {rule_tag} verification failed," f"expected: {expected}, got: {evaluation}"


register_params(
    test_eks_file_system_configuration,
    Parameters(
        ("rule_tag", "node_hostname", "expected"),
        [
            *eks_fs_tc.cis_eks_3_1_1.values(),
            *eks_fs_tc.cis_eks_3_1_2.values(),
            *eks_fs_tc.cis_eks_3_1_3.values(),
            *eks_fs_tc.cis_eks_3_1_4.values(),
        ],
        ids=[
            *eks_fs_tc.cis_eks_3_1_1.keys(),
            *eks_fs_tc.cis_eks_3_1_2.keys(),
            *eks_fs_tc.cis_eks_3_1_3.keys(),
            *eks_fs_tc.cis_eks_3_1_4.keys(),
        ],
    ),
)
