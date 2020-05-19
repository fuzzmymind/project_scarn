<template>
  <div class="landing-page">
    <h1>Welcome to Project Scarn</h1>
    <hr />
    <Message :text="loadText" />
  </div>
</template>

<script>
import Message from "@/components/Utility/Message";
export default {
  layout: "Default",
  components: {
    Message,
  },
  computed: {
    loadText() {
      return this.$store.getters.text;
    },
  },
};
</script>
<style scoped>
.landing-page {
  text-align: center;
}
</style>
