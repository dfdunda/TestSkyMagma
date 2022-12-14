"""
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

import unittest

import s1ap_types
from integ_tests.s1aptests import s1ap_wrapper


class TestAttachDetachSecurityAlgoEea1Eia1(unittest.TestCase):
    """Integration Test: TestAttachDetachSecurityAlgoEea1Eia1"""

    def setUp(self):
        """Initialize before test case execution"""
        self._s1ap_wrapper = s1ap_wrapper.TestWrapper()

    def tearDown(self):
        """Cleanup after test case execution"""
        self._s1ap_wrapper.cleanup()

    def test_attach_detach_security_algo_eea1_eia1(self):
        """
        Basic attach/detach test case with following security algorithms:
        Encryption algorithm type : Snow3G/EEA1
        Integrity algorithm type  : Snow3G/EIA1

        The preferred Encryption/Integrity algorythms to be selected by the
        network can be changed by manually modifying the order of preference of
        algorithms in mme.conf.template file
        """
        num_ues = 2
        detach_type = [
            s1ap_types.ueDetachType_t.UE_NORMAL_DETACH.value,
            s1ap_types.ueDetachType_t.UE_SWITCHOFF_DETACH.value,
        ]
        wait_for_s1 = [True, False]

        # Configure Encryption/Integrity algorithms for UE Network Capabilities
        config_list = []
        for i in range(num_ues):
            config_list.append(s1ap_types.ueConfig_t())
            config_list[i].ueNwCap_pr.pres = 1
            config_list[i].ueNwCap_pr.eea2_128 = 0
            config_list[i].ueNwCap_pr.eea1_128 = 1
            config_list[i].ueNwCap_pr.eea0 = 0
            config_list[i].ueNwCap_pr.eia2_128 = 0
            config_list[i].ueNwCap_pr.eia1_128 = 1
            config_list[i].ueNwCap_pr.eia0 = 0
        self._s1ap_wrapper.configUEDevice(num_ues, req_data=config_list)

        for i in range(num_ues):
            req = self._s1ap_wrapper.ue_req
            print(
                "************************* Running End to End attach for ",
                "UE id ",
                req.ue_id,
            )
            # Now actually complete the attach
            self._s1ap_wrapper._s1_util.attach(
                req.ue_id,
                s1ap_types.tfwCmd.UE_END_TO_END_ATTACH_REQUEST,
                s1ap_types.tfwCmd.UE_ATTACH_ACCEPT_IND,
                s1ap_types.ueAttachAccept_t,
            )

            # Wait on EMM Information from MME
            self._s1ap_wrapper._s1_util.receive_emm_info()
            print(
                "************************* Running UE detach for UE id ",
                req.ue_id,
            )
            # Now detach the UE
            self._s1ap_wrapper.s1_util.detach(
                req.ue_id,
                detach_type[i],
                wait_for_s1[i],
            )


if __name__ == "__main__":
    unittest.main()
