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

import time
import unittest

import s1ap_types
import s1ap_wrapper


class TestAttachMissingImsi(unittest.TestCase):

    def setUp(self):
        self._s1ap_wrapper = s1ap_wrapper.TestWrapper()

    def tearDown(self):
        self._s1ap_wrapper.cleanup()

    def test_attach_missing_imsi(self):
        """ Attaching with IMSI missing from subscriberd """
        self._s1ap_wrapper.configUEDevice(1)
        req = self._s1ap_wrapper.ue_req
        ue_id = req.ue_id
        print(
            "************************* Running End to End attach for ",
            "UE id ", ue_id,
        )
        self._s1ap_wrapper._sub_util.cleanup()
        # Now actually attempt the attach
        self._s1ap_wrapper._s1_util.attach(
            ue_id, s1ap_types.tfwCmd.UE_END_TO_END_ATTACH_REQUEST,
            s1ap_types.tfwCmd.UE_ATTACH_REJECT_IND,
            s1ap_types.ueAttachRejInd_t,
        )

        response = self._s1ap_wrapper.s1_util.get_response()
        assert response.msg_type == s1ap_types.tfwCmd.UE_CTX_REL_IND.value

        ue_id = 2
        print("************************* Adding IMSI entry for UE id ", ue_id)
        # Adding IMSI to subscriberdb
        self._s1ap_wrapper.configUEDevice_without_checking_gw_health(1)

        time.sleep(5)

        print(
            "************************* Rerunning End to End attach for ",
            "UE id ", ue_id,
        )
        # Now actually complete the attach
        self._s1ap_wrapper._s1_util.attach(
            ue_id, s1ap_types.tfwCmd.UE_END_TO_END_ATTACH_REQUEST,
            s1ap_types.tfwCmd.UE_ATTACH_ACCEPT_IND,
            s1ap_types.ueAttachAccept_t,
        )

        # Wait on EMM Information from MME
        self._s1ap_wrapper._s1_util.receive_emm_info()

        print("************************* Running UE detach for UE id ", ue_id)
        # Now detach the UE
        self._s1ap_wrapper.s1_util.detach(
            ue_id, s1ap_types.ueDetachType_t.UE_SWITCHOFF_DETACH.value, False,
        )


if __name__ == "__main__":
    unittest.main()
