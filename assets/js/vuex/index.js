import Vuex from 'vuex'

import user from './user';
import socket from './socket';

export default new Vuex.Store({
  modules: {
    user,
    socket
  }
})
