<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h2>Pizza iemand?</h2>
      <p v-for='pizza in pizzas' :key='pizza.id' >
        🍕 {{ pizza.name }}, 🏤 {{ pizza.author }}, ⭐ {{ pizza.rating }}
      </p>
  </div>
</template>

<script>
export default {
  name: 'Pizza World',
  props: {
    msg: String
  },
  data() {
    return {
      pizzas: {}
    }
  },
  created() {
    fetch('/api/v1/pizza')
        .then(result => result.json())
        .then(result => {
          this.pizzas = result
        })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
