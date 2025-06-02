import js from '@eslint/js';
import globals from 'globals';
import reactHooks from 'eslint-plugin-react-hooks';
import reactRefresh from 'eslint-plugin-react-refresh';
import tseslint from 'typescript-eslint';

export default tseslint.config(
  { ignores: ['dist'] }, // dist フォルダは対象外
  {
    extends: [js.configs.recommended, ...tseslint.configs.recommended], // 基本ルール
    files: ['**/*.{ts,tsx}'], // TypeScriptファイルのみ対象 (JavaScriptも含む場合は **/*.{js,jsx,ts,tsx})
    languageOptions: {
      ecmaVersion: 2020,
      globals: globals.browser,
    },
    // プラグインの設定
    plugins: {
      'react-hooks': reactHooks,
      'react-refresh': reactRefresh,
    },
    // 具体的なルール設定
    rules: {
      ...reactHooks.configs.recommended.rules, // React Hooksの推奨ルール
      'react-refresh/only-export-components': ['warn', { allowConstantExport: true }], // React Refreshのルール
      '@typescript-eslint/no-explicit-any': 'error', // anyの使用を禁止
    },
  }
);
