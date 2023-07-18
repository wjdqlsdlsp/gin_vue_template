<template>
  <div>
    <button @click="btnTest">Axios Test - Get</button>
  </div>

  <br>
  <div>
    <h2>{{ a }}</h2>
  </div>

  <div>
    <h2>{{ b }}</h2>
  </div>

  <div>
    <h2>{{ response_data }}</h2>
  </div>

  <br>

  <div>
    <div>
      id
      <input type="text" name="id" v-model="id" />
    </div>
    <div>
      password
      <input type="text" name="password" v-model="password" />
    </div>
    <div>
      <button @click="btnTest2">Axios Test - Post</button>
    </div>
  </div>

  <div>
  </div>

</template>

<script>
import axios from 'axios';

export default {
  data: () => ({
    b: '',
    a: '',
    c: '',
    response_data: '',
  }),
  methods: {
    btnTest() {
      axios.get('api/btnTest').then((response) => {
        this.a = response.data.a
        this.b = response.data.b
        this.c = response.data.c
          }).catch((error)=>{
          });
    },

    btnTest2() {
      axios.post('api/btnTest',
          {
            id: this.id,
            password: this.password
          })
          .then(res => {
            this.response_data = res.data.message
            console.log(res.data)
          });
    },
  }
}
</script>