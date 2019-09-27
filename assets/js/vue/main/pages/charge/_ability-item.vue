<template>
  <div class="ability-item">
    <span>{{ priv.privilege.name }}</span>
    <i class="el-icon-delete-solid" @click="deleteAbility"></i>
  </div>
</template>
<script>
import abilityDelete from './gql/mutation.ability-delete.gql';
import chargeQuery from './gql/query.charge.gql';

export default {
  name: 'ability-item',
  props: ['priv'],
  methods: {
    deleteAbility() {
      this.$apollo.mutate({
        mutation: abilityDelete,
        variables: { id: this.priv.id },
        update: (store, { data: { id } }) => {
          var opts = {
            query: chargeQuery,
            variables: this.$parent.variables
          };
          var data = store.readQuery(opts);
          var index = data.charge.abilities.findIndex(a => a.id === this.priv.id);
          data.charge.abilities.splice(index, 1);
          store.writeQuery({ ...opts, data });
        }
      });
    }
  }
};
</script>
