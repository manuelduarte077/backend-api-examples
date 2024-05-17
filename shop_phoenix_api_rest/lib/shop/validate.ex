defmodule Shop.Validate do
  @moduledoc """
  Provides functions for validating input parameters.
  """

  @doc """
  Returns {:ok, value} when field is present, or {:error, message} when field is nil.


  ## Examples
    iex> Shop.Validate.get_required(%{name: "John"}, :name)
    {:ok, "John"}

  """
  def get_required(params, field) do
    case params[field] do
      nil -> {:error, "Field #{field} is required"}
      value -> {:ok, value}
    end
  end

@doc """
    Validates if a given value is an integer.

    ## Examples

            iex> Shop.Validate.is_integer("123", :age)
            {:ok, 123}

            iex> Shop.Validate.is_integer("abc", :age)
            {:error, "age: Value is not an integer"}

    ## Parameters

        * `value` - The value to be validated.
        * `field` - The name of the field being validated.

    ## Returns

        * `{:ok, int}` - If the value is a valid integer.
        * `{:error, message}` - If the value is not an integer.

    """
    def is_integer(value, field) do
        case Integer.parse(value) do
            {int, ""} -> {:ok, int}
            _ -> {:error, "#{field}: Value is not an integer"}
        end
    end
end
