<script setup lang="ts">
import { ref, onMounted, watch, onUnmounted } from 'vue';
import { EditorState } from '@codemirror/state';
import { EditorView, keymap, lineNumbers, highlightActiveLine } from '@codemirror/view';
import { defaultKeymap } from "@codemirror/commands";
import { eclipse } from '@uiw/codemirror-theme-eclipse';
import { go } from '@codemirror/lang-go';

const props = defineProps<{
  modelValue: string;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();

const editorRef = ref<HTMLDivElement>();
let view: EditorView;

onMounted(() => {
  if (!editorRef.value) return;

  const updateListener = EditorView.updateListener.of((update) => {
    if (update.docChanged) {
      emit('update:modelValue', update.state.doc.toString());
    }
  });

  const state = EditorState.create({
    doc: props.modelValue,
    extensions: [
      go(),
      lineNumbers(),
      highlightActiveLine(),
      keymap.of(defaultKeymap),
      eclipse,
      updateListener,
    ],
  });

  view = new EditorView({
    state,
    parent: editorRef.value,
  });
});

watch(() => props.modelValue, (newValue) => {
  if (view && newValue !== view.state.doc.toString()) {
    view.dispatch({
      changes: { from: 0, to: view.state.doc.length, insert: newValue },
    });
  }
});

onUnmounted(() => {
  view?.destroy();
});
</script>

<template>
  <div ref="editorRef" class="editor-container" />
</template>

<style scoped>
.editor-container {
  border: 1px solid var(--border);
  border-radius: 8px;
  overflow: hidden;
  height: 250px;
  font-size: 14px;
}
.editor-container :deep(.cm-editor) {
  height: 100%;
}
</style>
