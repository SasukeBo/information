<template>
  <div
    class="stop-reason"
    v-loading="$apollo.queries.logs.loading"
    element-loading-background="unset"
  >
    <div v-if="!$apollo.queries.logs.loading && !logs.length" class="empty-reason">停机原因未知</div>
    <div v-for="(log, i) in logs" :key="'log_' + i" class="log-item">
      <span class="index">{{ i+1 }}.</span>
      <span class="content">{{ log.content }}</span>
    </div>
  </div>
</template>
<script>
import gql from 'graphql-tag';

export default {
  name: 'stop-reason',
  props: ['id'],
  apollo: {
    logs: {
      query: gql`
        query logStopReasonsGet($logID: Int!) {
          logs: logStopReasonsGet(logID: $logID) {
            content
          }
        }
      `,
      variables() {
        return {
          logID: this.id
        };
      }
    }
  },
  data() {
    return {
      logs: []
    };
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.stop-reason {
  min-height: 22px;

  .empty-reason {
    text-align: center;
    font-size: 0.9rem;
    color: $--color-font__gray;
  }

  .log-item {
    padding-bottom: 0.5rem;
    padding-left: 10px;
    color: $--color-theme__danger;

    &:last-child {
      padding-bottom: 0;
    }

    .index {
      padding: 0 1rem;
      color: $--color-font__gray;
    }
  }
}
</style>
