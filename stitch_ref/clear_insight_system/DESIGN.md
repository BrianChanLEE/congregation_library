---
name: Clear Insight System
colors:
  surface: '#f8f9fa'
  surface-dim: '#d9dadb'
  surface-bright: '#f8f9fa'
  surface-container-lowest: '#ffffff'
  surface-container-low: '#f3f4f5'
  surface-container: '#edeeef'
  surface-container-high: '#e7e8e9'
  surface-container-highest: '#e1e3e4'
  on-surface: '#191c1d'
  on-surface-variant: '#43474c'
  inverse-surface: '#2e3132'
  inverse-on-surface: '#f0f1f2'
  outline: '#74777d'
  outline-variant: '#c4c6cd'
  surface-tint: '#4e6073'
  primary: '#162839'
  on-primary: '#ffffff'
  primary-container: '#2c3e50'
  on-primary-container: '#96a9be'
  inverse-primary: '#b5c8df'
  secondary: '#006497'
  on-secondary: '#ffffff'
  secondary-container: '#77c2ff'
  on-secondary-container: '#004f79'
  tertiary: '#002e14'
  on-tertiary: '#ffffff'
  tertiary-container: '#004721'
  on-tertiary-container: '#3cbe6e'
  error: '#ba1a1a'
  on-error: '#ffffff'
  error-container: '#ffdad6'
  on-error-container: '#93000a'
  primary-fixed: '#d1e4fb'
  primary-fixed-dim: '#b5c8df'
  on-primary-fixed: '#091d2e'
  on-primary-fixed-variant: '#36485b'
  secondary-fixed: '#cce5ff'
  secondary-fixed-dim: '#92ccff'
  on-secondary-fixed: '#001e31'
  on-secondary-fixed-variant: '#004b73'
  tertiary-fixed: '#7efba4'
  tertiary-fixed-dim: '#61de8a'
  on-tertiary-fixed: '#00210c'
  on-tertiary-fixed-variant: '#005228'
  background: '#f8f9fa'
  on-background: '#191c1d'
  surface-variant: '#e1e3e4'
typography:
  headline-lg:
    fontFamily: Hanken Grotesk
    fontSize: 32px
    fontWeight: '700'
    lineHeight: '1.3'
    letterSpacing: -0.02em
  headline-md:
    fontFamily: Hanken Grotesk
    fontSize: 24px
    fontWeight: '600'
    lineHeight: '1.4'
    letterSpacing: -0.01em
  headline-sm:
    fontFamily: Hanken Grotesk
    fontSize: 18px
    fontWeight: '600'
    lineHeight: '1.4'
  body-lg:
    fontFamily: Noto Sans
    fontSize: 16px
    fontWeight: '400'
    lineHeight: '1.6'
  body-md:
    fontFamily: Noto Sans
    fontSize: 14px
    fontWeight: '400'
    lineHeight: '1.5'
  body-sm:
    fontFamily: Noto Sans
    fontSize: 13px
    fontWeight: '400'
    lineHeight: '1.5'
  label-md:
    fontFamily: Noto Sans
    fontSize: 12px
    fontWeight: '600'
    lineHeight: '1'
    letterSpacing: 0.02em
  headline-lg-mobile:
    fontFamily: Hanken Grotesk
    fontSize: 24px
    fontWeight: '700'
    lineHeight: '1.3'
rounded:
  sm: 0.125rem
  DEFAULT: 0.25rem
  md: 0.375rem
  lg: 0.5rem
  xl: 0.75rem
  full: 9999px
spacing:
  base: 4px
  xs: 4px
  sm: 8px
  md: 16px
  lg: 24px
  xl: 40px
  container-max: 1200px
  gutter: 16px
---

## Brand & Style

이 디자인 시스템은 고밀도 정보의 체계적인 전달과 깊은 몰입감을 최우선으로 합니다. JW Library의 사용성 철학을 계승하여, 복잡한 데이터를 논리적이고 정돈된 방식으로 시각화합니다.

**디자인 스타일: 정보 중심 미니멀리즘 (Information-Centric Minimalism)**
- **정밀함:** 불필요한 장식을 배제하고, 여백과 타이포그래피의 위계만으로 정보를 구분합니다.
- **신뢰성:** 안정적인 네이비 톤과 차분한 화이트 배경을 통해 도구로서의 신뢰감을 부여합니다.
- **고밀도:** 한 화면에 많은 정보를 담되, 시각적 소음(Noise)을 최소화하여 인지 부하를 낮춥니다.

## Colors

본 시스템은 청결함과 전문성을 강조하는 절제된 팔레트를 사용합니다.

- **Primary (#2C3E50):** 깊은 네이비 톤으로 헤더, 주요 텍스트, 브랜드 요소에 사용되어 무게감을 부여합니다.
- **Secondary (#2980B9):** 활기찬 블루 톤으로 버튼, 링크, 선택 상태 등 상호작용 요소에 적용합니다.
- **Background:** 메인 배경은 `#FFFFFF`를 사용하여 텍스트 가독성을 극대화하며, 섹션 구분이나 사이드바에는 `#F8F9FA`를 활용하여 층위를 만듭니다.
- **Status Badges:** 상태 표시를 위한 그린, 옐로우, 그레이는 명도를 높여 가독성을 확보한 파스텔 톤 기반의 원색을 사용합니다.

## Typography

한국어 환경에서 최상의 가독성을 제공하기 위해 `Noto Sans`를 본문 폰트로 채택하고, 영문 및 숫자 조합의 정밀함을 위해 `Hanken Grotesk`를 제목용으로 혼용합니다.

- **본문 가독성:** 장문의 텍스트 읽기를 고려하여 본문(`body-md`)의 행간을 1.5배 이상으로 설정합니다.
- **위계 강조:** 정보 밀도가 높은 화면에서는 폰트 크기 조절보다는 굵기(Weight) 변화와 컬러(Primary vs Secondary Text) 대비를 통해 우선순위를 구분합니다.
- **숫자 데이터:** 데이터 테이블 내 숫자는 자간을 일정하게 유지하여 수치 비교가 용이하도록 합니다.

## Layout & Spacing

4px 그리드 시스템을 기반으로 한 고밀도 레이아웃을 지향합니다.

- **Layout:** 데스크탑에서는 12컬럼 그리드를 사용하며, 최대 폭을 `1200px`로 제한하여 시선의 분산을 막습니다.
- **Density:** 컴포넌트 간 간격은 최소화(`sm`, `md`)하여 한 화면에 노출되는 정보량을 높이되, 그룹 간 간격은 확실하게(`lg`, `xl`) 벌려 구조적 명확성을 확보합니다.
- **Mobile Adaptation:** 모바일 환경에서는 1컬럼 플로우로 전환되며, 좌우 여백을 `16px`로 고정하여 가독 영역을 최대한 확보합니다. 상단 검색바는 스크롤 시에도 상단에 고정(Sticky)됩니다.

## Elevation & Depth

본 시스템은 그림자 효과를 최소화하고 **Tonal Layers(톤 층위)**를 통해 깊이감을 표현합니다.

- **Surface Levels:** 
  - Level 0 (Base): `#F8F9FA` 메인 배경.
  - Level 1 (Card/Content): `#FFFFFF` 배경 위에 얇은 경계선(`border: 1px solid #E9ECEF`)을 사용한 카드 영역.
- **Interaction:** 마우스 호버 시에만 아주 부드러운 앰비언트 섀도우(Blur 10px, Opacity 5%)를 적용하여 요소가 떠오르는 느낌을 줍니다.
- **Modals:** 모달 팝업은 다크 반투명 오버레이를 배경에 깔아 현재 문맥을 유지하면서도 포커스를 집중시킵니다.

## Shapes

신뢰감 있고 정돈된 인상을 위해 **Soft (0.25rem)** 라운딩을 기본으로 사용합니다.

- **Buttons & Inputs:** 4px(`0.25rem`)의 미세한 곡률을 적용하여 너무 딱딱하지 않으면서도 전문적인 인상을 유지합니다.
- **Cards & Modals:** 큰 컨테이너 요소는 8px(`0.5rem`)을 적용하여 내부 요소와의 시각적 위계를 분리합니다.
- **Chips/Badges:** 상태 표시용 배지는 텍스트를 감싸는 최소한의 높이에 맞춰 완전한 곡선(Pill-shaped)을 적용하여 기능 버튼과 형태적으로 구분합니다.

## Components

이 디자인 시스템의 컴포넌트는 작지만 명확한 사용성을 목표로 합니다.

- **고밀도 카드 (Compact Cards):** 패딩을 `12px~16px`로 제한하고, 상단에 배지를 배치하여 상태를 즉각 파악할 수 있게 합니다. 하단에는 액션 버튼 대신 아이콘 버튼을 배치하여 공간을 절약합니다.
- **상단 고정 검색바 (Sticky Search):** 배경색과 대비되는 화이트 필드에 그림자 없이 1px 보더를 사용합니다. 필터 탭은 검색바 바로 아래에 배치하여 검색 결과의 세부 조정을 돕습니다.
- **탭 형태 필터 (Segmented Tabs):** 선택된 탭은 네이비(#2C3E50) 배경에 화이트 텍스트로 표시하며, 선택되지 않은 탭은 고스트 스타일로 배경에 녹아들게 설계합니다.
- **데이터 테이블:** 행(Row) 간의 보더를 최소화하고 호버 시 하이라이트 컬러(`Secondary 5%`)를 적용합니다. 헤더는 `Primary` 컬러의 옅은 틴트를 배경으로 사용하여 데이터 영역과 구분합니다.
- **모달 팝업:** 제목 영역을 명확히 분리하고, 우측 상단에 '닫기' 아이콘, 우측 하단에 실행 버튼을 배치하는 표준 가이드를 준수합니다.