# License Policy

すべてのリポジトリに `LICENSE` ファイルを含める。ライセンスなしでの公開は禁止。

## License Selection

ライセンスは **Closed / AGPL・GPL・LGPL / MIT** の3分類から選択する。OSS だからといって必ずしも MIT ではない。

| プロジェクト種別 | 推奨ライセンス | 備考 |
|---|---|---|
| 研究・実験系（配布しない） | Closed | All Rights Reserved |
| 研究・実験系（配布する） | MIT | 制限なく利用できるよう MIT を基本とする |
| コンパイル配布系（Swift アプリ・Chrome 拡張など） | Closed または AGPL/GPL/LGPL | 将来の再ライセンスの余地を確保したい場合は準 CLA 対応を推奨（[topics/GITHUB_CONTRIBUTING.md](topics/GITHUB_CONTRIBUTING.md) 参照） |
| ソースコードで配布系（Alfred、Python など） | Closed（配布しない）または MIT | |
| 配布なし（Web app） | Closed | API ドキュメントのみ公開も選択肢 |
| 商用利用を制限したいドキュメント系 | CC BY-NC 4.0 など | ドキュメント系に限る |

ライセンス選択は case by case。上記は傾向の参考にとどめ、プロジェクトの状況に応じて判断すること。

## License Templates

### All Rights Reserved

クローズドプロジェクト向け：

```
Copyright (c) [YEAR] [AUTHOR]

All rights reserved.

This software and associated documentation files are proprietary and
confidential. Unauthorized copying, distribution, modification, or use
of this software, in whole or in part, is strictly prohibited without
the express written permission of the author.
```

### MIT（パーミッシブ）

ライセンス本文は公式ソースから取得すること：

- MIT: https://opensource.org/license/mit

`LICENSE` ファイルの先頭に `Copyright (c) [YEAR] [AUTHOR]` を追加した上で配置する。

### AGPL / GPL / LGPL（コピーレフト）

コピーレフトを採用する場合、ライセンス本文は公式ソースから取得すること：

- AGPL v3: https://www.gnu.org/licenses/agpl-3.0.txt
- GPL v3: https://www.gnu.org/licenses/gpl-3.0.txt
- LGPL v3: https://www.gnu.org/licenses/lgpl-3.0.txt

`LICENSE` ファイルの先頭に `Copyright (c) [YEAR] [AUTHOR]` を追加した上で配置する。

**選択の目安：**

| ライセンス | 用途 |
|---|---|
| AGPL v3 | Web サービス・SaaS 等、サーバサイドでの使用にもコピーレフトを及ぼしたい場合 |
| GPL v3 | スタンドアローンアプリ等、配布バイナリにコピーレフトを及ぼしたい場合 |
| LGPL v3 | ライブラリ等、リンクするコードへの影響を最小化したい場合 |

将来的な再ライセンスの余地を確保したい場合は [topics/GITHUB_CONTRIBUTING.md](topics/GITHUB_CONTRIBUTING.md) の準 CLA 設定を参照。
