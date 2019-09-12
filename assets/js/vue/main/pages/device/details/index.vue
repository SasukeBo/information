<template>
  <div class="device-details" v-if="!$apollo.queries.device.loading">
    <div class="global-card flex-content">
      <div class="device-details__left">
        <div class="param-chart">
          <value-chart v-if="!$apollo.queries.params.loading" :params="params"></value-chart>
        </div>
      </div>

      <div class="device-details__right">
        <device-status-card
          :device="device"
          :duration="duration"
          :status="status"
          :statusTag="statusTag"
        ></device-status-card>

        <base-card :device="device"></base-card>

        <register-card :device="device"></register-card>
      </div>
    </div>
  </div>
</template>
<script>
import deviceQuery from './gql/query.device-get.gql';
import paramsQuery from './gql/query.params.gql';
import deviceStatusSub from 'js/vue/main/pages/devices/gql/sub.device-status.gql';
import durationQuery from './gql/query.duration.gql';

import DeviceStatusCard from './_device-status-card';
import ValueChart from './value-chart';
import BaseCard from './_base-card';
import RegisterCard from './_register-card';

export default {
  name: 'device-details',
  props: ['uuid'],
  components: {
    BaseCard,
    ValueChart,
    RegisterCard,
    DeviceStatusCard
  },
  apollo: {
    device: {
      query: deviceQuery,
      variables() {
        return { uuid: this.uuid };
      },
      fetchPolicy: 'network-only'
    },
    params: {
      query: paramsQuery,
      variables() {
        return { deviceUUID: this.uuid };
      }
    },
    $subscribe: {
      deviceUpdate: {
        query: deviceStatusSub,
        variables() {
          return {
            t: `dsl:${this.device.id}`
          };
        },
        result({ data }) {
          this.device.status = data.deviceUpdate.status;
          this.device.remoteIP = data.deviceUpdate.remoteIP;
        }
      }
    }
  },
  data() {
    return {
      duration: '0时0分0秒',
      device: {},
      params: [],
      statusTag: '',
      statusMap: {
        prod: '生产中',
        offline: '离线',
        stop: '停机'
      }
    };
  },
  watch: {
    device(newVal) {
      this.statusTag = newVal.status;
      this.$apollo
        .query({
          query: durationQuery,
          variables: {
            deviceID: newVal.id,
            status: 'prod'
          },
          fetchPolicy: 'network-only'
        })
        .then(({ data }) => {
          this.duration = data.duration;
        })
        .catch(e => {
          console.log(e);
        });
    }
  },
  computed: {
    status() {
      return this.statusMap[this.statusTag];
    }
  }
};
</script>
