/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the terms found in the LICENSE file in the root of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */

#include "lte/gateway/c/core/oai/tasks/nas/emm/msg/GutiReallocationComplete.hpp"

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif
#include "lte/gateway/c/core/oai/common/TLVEncoder.h"
#include "lte/gateway/c/core/oai/common/TLVDecoder.h"
#ifdef __cplusplus
}
#endif

int decode_guti_reallocation_complete(
    guti_reallocation_complete_msg* guti_reallocation_complete, uint8_t* buffer,
    uint32_t len) {
  uint32_t decoded = 0;

  // Check if we got a NULL pointer and if buffer length is >= minimum length
  // expected for the message.
  CHECK_PDU_POINTER_AND_LENGTH_DECODER(
      buffer, GUTI_REALLOCATION_COMPLETE_MINIMUM_LENGTH, len);
  /*
   * Decoding mandatory fields
   */
  return decoded;
}

int encode_guti_reallocation_complete(
    guti_reallocation_complete_msg* guti_reallocation_complete, uint8_t* buffer,
    uint32_t len) {
  int encoded = 0;

  /*
   * Checking IEI and pointer
   */
  CHECK_PDU_POINTER_AND_LENGTH_ENCODER(
      buffer, GUTI_REALLOCATION_COMPLETE_MINIMUM_LENGTH, len);
  return encoded;
}
