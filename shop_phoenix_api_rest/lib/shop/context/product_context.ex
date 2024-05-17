defmodule Shop.ProductContext do
  @moduledoc false

  alias Shop.Core.Product

  def all(_params) do
    case Product.Api.all() do
      {:ok, products} -> {:ok, products |> Enum.map(&Product.Api.json!(&1, :public))}
      error -> error
    end
  end

  def show(params) do
    with {:ok, id} <- Shop.Validate.get_required(params, "id"),
         {:ok, _id} <- Shop.Validate.is_integer(id, "id"),
         {:ok, product} <- Product.Api.get(id) do
      {:ok, Product.Api.json!(product, :public)}
    else
      {:error, msg} -> {:error, msg}
    end
  end

  def create(params) do
    case Product.Api.insert(params) do
      {:ok, product} -> {:ok, Product.Api.json!(product, :public)}
      {:error, %Ecto.Changeset{} = changeset} -> {:error, parse_errors(changeset)}
    end
  end

  def update(params) do
    with {:ok, id} <- Shop.Validate.get_required(params, "id"),
         {:ok, _id} <- Shop.Validate.is_integer(id, "id"),
         {:ok, product} <- Product.Api.get(id),
         {:ok, updated_product} <- Product.Api.update(product, params) do
      {:ok, updated_product |> Product.Api.json!(:public)}
    else
      {:error, %Ecto.Changeset{} = changeset} ->
        {:error, parse_errors(changeset)}

      {:error, msg} ->
        {:error, msg}
    end
  end

  def delete(params) do
    with {:ok, id} <- Shop.Validate.get_required(params, "id"),
         {:ok, _id} <- Shop.Validate.is_integer(id, "id"),
         {:ok, product} <- Product.Api.delete(id) do
      {:ok, product |> Product.Api.json!(:public)}
    else
      {:error, msg} -> {:error, msg}
    end
  end

  defp parse_errors(changeset) do
    Enum.map(changeset.errors, fn
      {key, {message, _}} -> "#{key}: #{message}"
    end)
  end
end
