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


import ctypes
import time
import unittest

import s1ap_types
import s1ap_wrapper


class TestImeiRestrictionSmc(unittest.TestCase):
    def setUp(self):
        self._s1ap_wrapper = s1ap_wrapper.TestWrapper()

    def tearDown(self):
        self._s1ap_wrapper.cleanup()

    def test_imei_restriction_smc(self):
        """
        This TC does the following:
        1. Send security mode complete message with a blocked imeisv
           which is same as the one configured in mme.conf.template
        2. Verify that attach reject is received with cause(5)
           IMEI_NOT_ACCEPTED
        3. Attach a 2nd UE with an allowed IMEISV by invoking attach utility
           function
        4. Detach the UE

        If this TC is executed individually run
        test_modify_mme_config_for_sanity.py to add
        { IMEI="9900048235103723"}
        under the BLOCKED_IMEI_LIST in mme.conf.template.

        After execution of this TC run test_restore_mme_config_after_sanity.py
        to restore the old mme.conf.template.
        """
        num_ues = 2
        ue_ids = []
        self._s1ap_wrapper.configUEDevice(num_ues)
        for _ in range(num_ues):
            req = self._s1ap_wrapper.ue_req
            ue_ids.append(req.ue_id)
        # Send Attach Request
        attach_req = s1ap_types.ueAttachRequest_t()
        sec_ctxt = s1ap_types.TFW_CREATE_NEW_SECURITY_CONTEXT
        id_type = s1ap_types.TFW_MID_TYPE_IMSI
        eps_type = s1ap_types.TFW_EPS_ATTACH_TYPE_EPS_ATTACH
        pdn_type = s1ap_types.pdn_Type()
        pdn_type.pres = True
        pdn_type.pdn_type = 1
        attach_req.ue_Id = ue_ids[0]
        attach_req.mIdType = id_type
        attach_req.epsAttachType = eps_type
        attach_req.useOldSecCtxt = sec_ctxt
        attach_req.pdnType_pr = pdn_type

        self._s1ap_wrapper._s1_util.issue_cmd(
            s1ap_types.tfwCmd.UE_ATTACH_REQUEST, attach_req,
        )
        print(
            "********************** Sent attach req for UE id ", ue_ids[0],
        )

        response = self._s1ap_wrapper.s1_util.get_response()
        assert response.msg_type == s1ap_types.tfwCmd.UE_AUTH_REQ_IND.value
        auth_req = response.cast(s1ap_types.ueAuthReqInd_t)
        print(
            "********************** Received auth req for UE id",
            auth_req.ue_Id,
        )

        # Send Authentication Response
        auth_res = s1ap_types.ueAuthResp_t()
        auth_res.ue_Id = ue_ids[0]
        sqnRecvd = s1ap_types.ueSqnRcvd_t()
        sqnRecvd.pres = 0
        auth_res.sqnRcvd = sqnRecvd
        self._s1ap_wrapper._s1_util.issue_cmd(
            s1ap_types.tfwCmd.UE_AUTH_RESP, auth_res,
        )
        print("********************** Sent auth rsp for UE id", ue_ids[0])

        response = self._s1ap_wrapper.s1_util.get_response()
        assert response.msg_type == s1ap_types.tfwCmd.UE_SEC_MOD_CMD_IND.value
        sec_mod_cmd = response.cast(s1ap_types.ueSecModeCmdInd_t)
        print(
            "********************** Received security mode cmd for UE id",
            sec_mod_cmd.ue_Id,
        )

        # Send Security Mode Complete
        sec_mode_complete = s1ap_types.ueSecModeComplete_t()
        sec_mode_complete.ue_Id = ue_ids[0]
        sec_mode_complete.imeisv_pres = True
        imeisv = "9900048235103723"
        # Check if the len of imeisv exceeds 16
        assert len(imeisv) <= 16
        for i in range(0, len(imeisv)):
            sec_mode_complete.imeisv[i] = ctypes.c_ubyte(int(imeisv[i]))
        self._s1ap_wrapper._s1_util.issue_cmd(
            s1ap_types.tfwCmd.UE_SEC_MOD_COMPLETE, sec_mode_complete,
        )
        print(
            "********************** Sent security mode complete for UE id",
            ue_ids[0],
        )

        # Receive Attach Reject
        response = self._s1ap_wrapper.s1_util.get_response()
        assert response.msg_type == s1ap_types.tfwCmd.UE_ATTACH_REJECT_IND.value
        attach_rej = response.cast(s1ap_types.ueAttachRejInd_t)
        print(
            "********************** Received attach reject for UE id %d"
            " with emm cause %d" % (attach_rej.ue_Id, attach_rej.cause),
        )

        # Verify cause
        assert attach_rej.cause == s1ap_types.TFW_EMM_CAUSE_IMEI_NOT_ACCEPTED

        # UE Context release
        response = self._s1ap_wrapper.s1_util.get_response()
        assert response.msg_type == s1ap_types.tfwCmd.UE_CTX_REL_IND.value

        ue_context_rel = response.cast(s1ap_types.ueCntxtRelReq_t)
        print(
            "********************** Received UE_CTX_REL_IND for UE id ",
            ue_context_rel.ue_Id,
        )

        # Attach the 2nd UE
        print(
            "************ Running end-end attach for UE id ************",
            ue_ids[1],
        )
        self._s1ap_wrapper._s1_util.attach(
            ue_ids[1],
            s1ap_types.tfwCmd.UE_END_TO_END_ATTACH_REQUEST,
            s1ap_types.tfwCmd.UE_ATTACH_ACCEPT_IND,
            s1ap_types.ueAttachAccept_t,
        )

        # Wait on EMM Information from MME
        self._s1ap_wrapper._s1_util.receive_emm_info()

        print("********************** Sleeping for 2 seconds")
        time.sleep(2)
        print("********************** Running UE detach for UE id ", ue_ids[1])
        # Now detach the UE
        self._s1ap_wrapper.s1_util.detach(
            ue_ids[1], s1ap_types.ueDetachType_t.UE_NORMAL_DETACH.value,
        )


if __name__ == "__main__":
    unittest.main()
