<template>
  <div class="alert alert-secondary" role="alert">
    <p>{{ text }}</p>
  </div>
</template>

<script>
export default {
  name: "Message",
  props: {
    text: {
      type: String,
      required: true
    }
  }
};
</script>
